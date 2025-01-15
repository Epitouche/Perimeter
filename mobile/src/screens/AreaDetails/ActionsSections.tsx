import React, { useState, useContext, useEffect } from 'react';
import { View, Text, TouchableOpacity, Modal, TextInput } from 'react-native';
import { SvgFromUri } from 'react-native-svg';
import { styles } from './StylesAreaDetails';
import { AppContext } from '../../context/AppContext';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

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
    const initialActionOptions = Object.entries(area.action_option).reduce(
      (acc, [name, value]) => {
        acc[name] = value;
        return acc;
      },
      {} as { [key: string]: any },
    );
    setSelectedActionOptions(initialActionOptions);
  }, [area.action_option]);

  const handleActionOptionChange = (key: string, value: any, type: any) => {
    setSelectedActionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

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
          <View>
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
                    keyboardType="default" // Adjust as needed
                    accessibilityHint={`Input for ${key}`}
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
                accessibilityHint="Saves the modified action options">
                <View style={styles.saveButton}>
                  <Text style={[{ color: 'white' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setIsActionModalVisible(false)}
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
