import React, { useState, useContext } from 'react';
import { Switch, View, Text, TouchableOpacity, Modal, TextInput } from 'react-native';
import { SvgFromUri } from 'react-native-svg';
import { styles } from './StylesAreaDetails';
import { AppContext } from '../../context/AppContext';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaSections = ({ navigation, route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [refreshRate, setRefreshRate] = useState<number>();
  const [isAreaModalVisible, setIsAreaModalVisible] = useState(false);
  const [isEnabled, setIsEnabled] = useState(area.enable);

  const handleSaveArea = async () => {
    console.log(title, description, refreshRate);
    const newArea = {
      ...area,
      title: title,
      description: description,
      refresh_rate: refreshRate,
    };
    console.log(newArea);
    try {
      const response = await fetch(`http://${ipAddress}:8080/api/v1/area`, {
        method: 'PUT',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(newArea),
      });
      console.log('body:', JSON.stringify(newArea));
      console.log(response);
      if (response.ok) {
        console.log('Area updated successfully');
      }
    } catch (error) {
      console.error('Error update area:', error);
    }
    setIsAreaModalVisible(false);
  };

  const deleteArea = async () => {
    try {
      const response = await fetch(`http://${ipAddress}:8080/api/v1/area/`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ id: area.id }),
      });
      if (response.ok) {
        console.log('Area deleted successfully');
        navigation.navigate('AreaView');
      }
    } catch (error) {
      console.error('Error deleting area:', error);
    }
  };

  const handleAreaStatus = async (value: boolean) => {
    const newArea = {
      ...area,
      enable: value,
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
  }

  return (
    <View>
      <View
        style={[
          styles.subContainer,
          { borderColor: 'black', borderWidth: 1, borderRadius: 10 },
        ]}>
        <View
          style={[{ flexDirection: 'row', justifyContent: 'space-between' }]}>
          <View style={[{ flexDirection: 'column' }]}>
            <View style={[styles.detailContainer, { flexDirection: 'column' }]}>
              <Text style={[styles.label, { color: 'black' }]}>title:</Text>
              <Text style={[styles.value, { color: 'black' }]}>
                {title == '' ? area.title : title}
              </Text>
            </View>
            <View style={[styles.detailContainer, { flexDirection: 'column' }]}>
              <Text style={[styles.label, { color: 'black' }]}>
                Description:
              </Text>
              <Text style={[styles.value, { color: 'black' }]}>
                {description == '' ? area.description : description}
              </Text>
            </View>
            <View style={styles.detailContainer}>
              <Text style={[styles.label, { color: 'black' }]}>
                Refresh rate:
              </Text>
              <Text style={[styles.value, { color: 'black' }]}>
                {refreshRate == undefined ? area.refresh_rate : refreshRate}
              </Text>
            </View>
          </View>
          <View
            style={{
              alignContent: 'center',
              justifyContent: 'space-between',
            }}>
            <TouchableOpacity
              onPress={() => deleteArea()}
              accessibilityLabel="Delete Area Button"
              accessibilityHint="Deletes the current area">
              <SvgFromUri
                uri={'https://api.iconify.design/mdi:delete.svg'}
                width={50}
                height={50}
                color={'#E60000'}
              />
            </TouchableOpacity>
            <Switch
              value={isEnabled}
              onValueChange={(value) => {setIsEnabled(value), handleAreaStatus(value)}}
              trackColor={{ false: "#E60000", true: "#1DC000" }}
              thumbColor={isEnabled ? "#000000" : "#000000"}
            />
            <TouchableOpacity
              onPress={() => setIsAreaModalVisible(true)}
              accessibilityLabel="Edit Area Button"
              accessibilityHint="Opens a modal to edit the area details">
              <SvgFromUri
                uri={'https://api.iconify.design/mdi:pencil-circle-outline.svg'}
                width={50}
                height={50}
                color={'black'}
              />
            </TouchableOpacity>
          </View>
        </View>
      </View>

      {/* Area Modal */}
      <Modal
        visible={isAreaModalVisible}
        transparent={true}
        animationType="slide">
        <View style={styles.modalContainer}>
          <View
            style={[
              styles.modalContent,
              {
                backgroundColor: 'white',
                borderColor: 'black',
                borderWidth: 1,
                borderRadius: 10,
              },
            ]}>
            <Text style={[styles.modalHeader, { color: 'black' }]}>
              Modify Area
            </Text>
            <View style={[{ flexDirection: 'column' }]}>
              <View style={[styles.optionRow]}>
                <Text style={[styles.optionLabel, { color: 'black' }]}>
                  Title
                </Text>
                <TextInput
                  style={styles.optionInput}
                  value={title}
                  onChangeText={text => setTitle(text)}
                  keyboardType="default" // Adjust as needed
                  accessibilityLabel="Title Input"
                  accessibilityHint="Input field for the area title"
                />
              </View>
              <View style={[styles.optionRow]}>
                <Text style={[styles.optionLabel, { color: 'black' }]}>
                  Description
                </Text>
                <TextInput
                  style={styles.optionInput}
                  value={description}
                  onChangeText={text => setDescription(text)}
                  keyboardType="default" // Adjust as needed
                  accessibilityLabel="Description Input"
                  accessibilityHint="Input field for the area description"
                />
              </View>
              <View style={[styles.optionRow]}>
                <Text style={[styles.optionLabel, { color: 'black' }]}>
                  Refresh rate
                </Text>
                <TextInput
                  style={[styles.optionInput]}
                  value={refreshRate ? String(refreshRate) : ''}
                  onChangeText={text => setRefreshRate(Number(text))}
                  keyboardType="numeric" // Adjust as needed
                  accessibilityLabel="Refresh Rate Input"
                  accessibilityHint="Input field for the area refresh rate"
                />
              </View>
            </View>
            <View
              style={{
                flexDirection: 'row',
                justifyContent: 'space-between',
              }}>
              <TouchableOpacity
                onPress={handleSaveArea}
                accessibilityLabel="Save Button"
                accessibilityHint="Saves the changes made to the area">
                <View style={[styles.saveButton, { borderColor: 'black' }]}>
                  <Text style={[{ color: 'black' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => {
                  setIsAreaModalVisible(false);
                  setDescription(area.description);
                  setTitle(area.title);
                  setRefreshRate(area.refresh_rate);
                }}
                accessibilityLabel="Cancel Button"
                accessibilityHint="Cancels the changes and closes the modal">
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

export default AreaSections;
