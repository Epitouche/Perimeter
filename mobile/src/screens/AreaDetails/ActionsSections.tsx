import React, { useState, useContext } from 'react';
import { View, Text, TouchableOpacity, Modal, TextInput } from 'react-native';
import { SvgFromUri } from 'react-native-svg';
import { styles } from './StylesAreaDetails';
import { AppContext } from '../../context/AppContext';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

/**
 * Component for displaying and modifying action options for a specific area.
 *
 * @param {Object} props - The component props.
 * @param {Object} props.route - The route object containing navigation parameters.
 * @param {Object} props.route.params - The parameters passed to the route.
 * @param {Object} props.route.params.area - The area object containing action options and service details.
 *
 * @returns {JSX.Element} The rendered component.
 *
 * @component
 * @example
 * <ActionsSections route={route} />
 *
 * @typedef {Object} Props
 * @property {Object} route - The route object containing navigation parameters.
 * @property {Object} route.params - The parameters passed to the route.
 * @property {Object} route.params.area - The area object containing action options and service details.
 */
const ActionsSections = ({ route }: Props) => {
  const { area } = route.params;
  const [isActionModalVisible, setIsActionModalVisible] = useState(false);
  const [selectedActionOptions, setSelectedActionOptions] = useState<{
    [key: string]: any;
  }>({});

  const { ipAddress, token } = useContext(AppContext);
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [refreshRate, setRefreshRate] = useState<number>();

  React.useEffect(() => {
    /**
     * Initializes the action options for the area by converting the entries of the `area.action_option` object
     * into a new object where each key-value pair is preserved.
     *
     * @param {Object} area - The area object containing action options.
     * @param {Object} area.action_option - The action options of the area.
     * @returns {Object} The initialized action options object.
     */
    const initialActionOptions = Object.entries(area.action_option).reduce(
      (acc, [name, value]) => {
        acc[name] = value;
        return acc;
      },
      {} as { [key: string]: any },
    );

    setSelectedActionOptions(initialActionOptions);
  }, [area.action_option]);

  /**
   * Handles the change of an action option.
   *
   * @param {string} key - The key of the action option to change.
   * @param {any} value - The new value of the action option.
   * @param {any} type - The type of the action option, used to determine if the value should be converted to a number.
   */
  const handleActionOptionChange = (key: string, value: any, type: any) => {
    setSelectedActionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

  /**
   * Handles the save action for updating an area.
   *
   * This function constructs a new area object with the selected action options,
   * sends a PUT request to update the area on the server, and updates the local
   * state with the response data if the request is successful.
   *
   * @async
   * @function handleSaveAction
   *
   * @returns {Promise<void>} A promise that resolves when the save action is complete.
   *
   * @throws Will log an error message if the update request fails.
   */
  const handleSaveAction = async () => {
    const newArea = {
      ...area,
      action_option: selectedActionOptions,
    };
    try {
      const response = await fetch(`http://${ipAddress}:8080/api/v1/area`, {
        method: 'PUT',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(newArea),
      });
      console.log(response);
      if (response.ok) {
        const body = await response.json();
        setDescription(body.description);
        setTitle(body.title);
        setRefreshRate(body.refresh_rate);
        console.log('Area updated successfully');
      }
    } catch (error) {
      console.error('Error update area:', error);
    }
    setIsActionModalVisible(false);
  };

  return (
    <View>
      <View
        style={[
          styles.subContainer,
          { backgroundColor: area.action.service.color },
        ]}>
        <View style={styles.ActionReactionHeader}>
          <Text style={styles.label}>Action</Text>
        </View>
        <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
          <View style={{ maxWidth: '60%' }}>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Service:</Text>
              <Text style={styles.value}>{area.action.service.name}</Text>
            </View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Options:</Text>
              <Text style={styles.value}>
                {Object.entries(selectedActionOptions).map(
                  ([key, value]) => `${key}: ${value} `,
                )}
              </Text>
            </View>
          </View>
          <TouchableOpacity
            onPress={() => setIsActionModalVisible(true)}
            accessibilityLabel="Modify Action"
            accessibilityHint="Opens a modal to modify the action options">
            <SvgFromUri
              uri={'https://api.iconify.design/mdi:pencil-circle-outline.svg'}
              width={50}
              height={50}
              color={'white'}
            />
          </TouchableOpacity>
        </View>
      </View>

      <Modal
        visible={isActionModalVisible}
        transparent={true}
        animationType="slide">
        <View style={styles.modalContainer}>
          <View
            style={[
              styles.modalContent,
              { backgroundColor: area.action.service.color },
            ]}>
            <Text style={styles.modalHeader}>Modify Action</Text>
            <View style={[{ flexDirection: 'column' }]}>
              {Object.keys(selectedActionOptions).map(key => (
                <View key={key} style={styles.optionRow}>
                  <Text style={styles.optionLabel}>{key}</Text>
                  <TextInput
                    style={styles.optionInput}
                    value={String(selectedActionOptions[key])}
                    onChangeText={text =>
                      handleActionOptionChange(
                        key,
                        text,
                        typeof selectedActionOptions[key],
                      )
                    }
                    keyboardType={`${
                      typeof selectedActionOptions[key] === 'number'
                        ? 'numeric'
                        : 'default'
                    }`}
                    accessibilityLabel="Action Option Input"
                    accessibilityHint={`Input for the ${key} option`}
                  />
                </View>
              ))}
            </View>
            <View
              style={{
                flexDirection: 'row',
                justifyContent: 'space-between',
              }}>
              <TouchableOpacity
                onPress={handleSaveAction}
                accessibilityLabel="Save Action"
                accessibilityHint="Saves the modified action options">
                <View style={styles.saveButton}>
                  <Text style={[{ color: 'white' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setIsActionModalVisible(false)}
                accessibilityLabel="Cancel Action"
                accessibilityHint="Closes the modal without saving changes">
                <View style={styles.cancelButton}>
                  <Text style={[{ color: 'red' }, { fontSize: 16 }]}>
                    Cancel
                  </Text>
                </View>
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </Modal>
    </View>
  );
};

export default ActionsSections;
