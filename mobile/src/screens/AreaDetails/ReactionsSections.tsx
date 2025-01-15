import React, { useState, useContext } from 'react';
import { View, Text, TouchableOpacity, Modal, TextInput } from 'react-native';
import { SvgFromUri } from 'react-native-svg';
import { styles } from './StylesAreaDetails';
import { AppContext } from '../../context/AppContext';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const ReactionsSections = ({ route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [isReactionModalVisible, setIsReactionModalVisible] = useState(false);
  const [selectedReactionOptions, setSelectedReactionOptions] = useState<{
    [key: string]: any;
  }>({});
  console.log('Area refreshrate:', area.action_refresh_rate);

  const handleReactionOptionChange = (key: string, value: any, type: any) => {
    setSelectedReactionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

  const handleSaveReaction = async () => {
    const newArea = {
      ...area,
      reaction_option: selectedReactionOptions,
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
        console.log('Area updated successfully');
      }
    } catch (error) {
      console.error('Error update area:', error);
    }
    setIsReactionModalVisible(false);
  };

  for (const option of Object.entries(area.reaction_option).map(
    ([name, value]) => ({ name, value }),
  )) {
    selectedReactionOptions[option.name] = option.value;
  }

  return (
    <View>
      <View
        style={[
          styles.subContainer,
          { backgroundColor: area.reaction.service.color },
        ]}>
        <View style={styles.ActionReactionHeader}>
          <Text
            style={styles.label}
            accessibilityLabel="Reaction Label"
            accessibilityHint="Indicates the reaction section">
            Reaction
          </Text>
        </View>
        <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
          <View>
            <View style={styles.detailContainer}>
              <Text
                style={styles.label}
                accessibilityLabel="Service Label"
                accessibilityHint="Label for the service name">
                Service:
              </Text>
              <Text
                style={styles.value}
                accessibilityLabel="Service Name"
                accessibilityHint="Name of the service">
                {area.reaction.service.name}
              </Text>
            </View>
            <View style={styles.detailContainer}>
              <Text
                style={styles.label}
                accessibilityLabel="Options Label"
                accessibilityHint="Label for the reaction options">
                Options:
              </Text>
              <Text
                style={styles.value}
                accessibilityLabel="Options Value"
                accessibilityHint="Values of the reaction options">
                {Object.entries(selectedReactionOptions).map(
                  ([key, value]) => `${key}: ${value} `,
                )}
              </Text>
            </View>
          </View>
          <TouchableOpacity
            onPress={() => setIsReactionModalVisible(true)}
            accessibilityLabel="Edit Reaction Button"
            accessibilityHint="Opens the modal to edit reaction options">
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
        visible={isReactionModalVisible}
        transparent={true}
        animationType="slide"
        accessibilityLabel="Reaction Modal"
        accessibilityHint="Modal to modify reaction options">
        <View style={styles.modalContainer}>
          <View
            style={[
              styles.modalContent,
              { backgroundColor: area.reaction.service.color },
            ]}>
            <Text
              style={styles.modalHeader}
              accessibilityLabel="Modify Reaction Header"
              accessibilityHint="Header for the modify reaction modal">
              Modify Reaction
            </Text>
            <View style={[{ flexDirection: 'column' }]}>
              {Object.keys(selectedReactionOptions).map(key => (
                <View key={key} style={styles.optionRow}>
                  <Text
                    style={styles.optionLabel}
                    accessibilityLabel={`${key} Label`}
                    accessibilityHint={`Label for the ${key} option`}>
                    {key}
                  </Text>
                  <TextInput
                    style={styles.optionInput}
                    value={String(selectedReactionOptions[key])}
                    onChangeText={text =>
                      handleReactionOptionChange(
                        key,
                        text,
                        typeof selectedReactionOptions[key],
                      )
                    }
                    keyboardType="default" // Adjust as needed
                    accessibilityLabel={`${key} Input`}
                    accessibilityHint={`Input field for the ${key} option`}
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
                onPress={handleSaveReaction}
                accessibilityLabel="Save Reaction Button"
                accessibilityHint="Saves the modified reaction options">
                <View style={styles.saveButton}>
                  <Text style={[{ color: 'white' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setIsReactionModalVisible(false)}
                accessibilityLabel="Cancel Reaction Button"
                accessibilityHint="Cancels the modification of reaction options">
                <View style={styles.cancelButton}>
                  <Text style={[{ color: '#E60000' }, { fontSize: 16 }]}>
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

export default ReactionsSections;
