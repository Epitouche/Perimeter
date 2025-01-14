import React, { useState, useContext } from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  Modal,
  TextInput,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';
import { AppContext } from '../context/AppContext';
import { SvgFromUri } from 'react-native-svg';
import BottomNavBar from './NavBar';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetailsScreen = ({ navigation, route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [isActionModalVisible, setIsActionModalVisible] = useState(false);
  const [isReactionModalVisible, setIsReactionModalVisible] = useState(false);
  const [isAreaModalVisible, setIsAreaModalVisible] = useState(false);
  const [selectedActionOptions, setSelectedActionOptions] = useState<{
    [key: string]: any;
  }>({});
  const [selectedReactionOptions, setSelectedReactionOptions] = useState<{
    [key: string]: any;
  }>({});

  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [refreshRate, setRefreshRate] = useState<number>();

  const handleActionOptionChange = (key: string, value: any, type: any) => {
    setSelectedActionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

  const handleReactionOptionChange = (key: string, value: any, type: any) => {
    setSelectedReactionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

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

  for (const option of Object.entries(area.reaction_option).map(
    ([name, value]) => ({ name, value }),
  )) {
    selectedReactionOptions[option.name] = option.value;
  }

  const renderModal = (
    isVisible: boolean,
    setIsVisible: React.Dispatch<React.SetStateAction<boolean>>,
    options: { [key: string]: any },
    handleOptionChange: (key: string, value: any, type: any) => void,
    handleSave: () => void,
    title: string,
    backgroundColor: string,
  ) => (
    <Modal visible={isVisible} transparent={true} animationType="slide">
      <View style={styles.modalContainer}>
        <View style={[styles.modalContent, { backgroundColor }]}>
          <Text style={styles.modalHeader}>{title}</Text>
          <View style={[{ flexDirection: 'column' }]}>
            {Object.keys(options).map(key => (
              <View key={key} style={styles.optionRow}>
                <Text style={styles.optionLabel}>{key}</Text>
                <TextInput
                  style={styles.optionInput}
                  value={String(options[key])}
                  onChange={event =>
                    handleOptionChange(
                      key,
                      event.nativeEvent.text,
                      typeof options[key],
                    )
                  }
                  keyboardType={
                    typeof options[key] === 'number' ? 'numeric' : 'default'
                  }
                />
              </View>
            ))}
          </View>
          <View
            style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
            <TouchableOpacity onPress={handleSave}>
              <View style={styles.saveButton}>
                <Text style={[{ color: 'white' }, { fontSize: 16 }]}>Save</Text>
              </View>
            </TouchableOpacity>
            <TouchableOpacity onPress={() => setIsVisible(false)}>
              <View style={styles.cancelButton}>
                <Text style={[{ color: 'red' }, { fontSize: 16 }]}>Cancel</Text>
              </View>
            </TouchableOpacity>
          </View>
        </View>
      </View>
    </Modal>
  );

  const getSections = (type: string) => {
    return (
      <View
        style={[
          styles.subContainer,
          {
            backgroundColor:
              type == 'action'
                ? area.action.service.color
                : area.reaction.service.color,
          },
        ]}>
        <View style={styles.ActionReactionHeader}>
          <Text style={styles.label}>
            {type == 'action' ? 'Action' : 'Reaction'}
          </Text>
        </View>
        <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
          <View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Service:</Text>
              <Text style={styles.value}>
                {type == 'action'
                  ? area.action.service.name
                  : area.reaction.service.name}
              </Text>
            </View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Options:</Text>
              <Text style={styles.value}>
                {Object.entries(
                  type === 'action'
                    ? selectedActionOptions
                    : selectedReactionOptions,
                ).map(([key, value]) => `${key}: ${value} `)}
              </Text>
            </View>
          </View>
          <TouchableOpacity
            onPress={() => {
              type == 'action'
                ? setIsActionModalVisible(true)
                : setIsReactionModalVisible(true);
            }}>
            <SvgFromUri
              uri={'https://api.iconify.design/mdi:pencil-circle-outline.svg'}
              width={50}
              height={50}
              color={'white'}
            />
          </TouchableOpacity>
        </View>
      </View>
    );
  };

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Area Details</Text>

      {/* Area Section */}
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
          <View style={{ alignContent: 'center' }}>
            <TouchableOpacity onPress={() => setIsAreaModalVisible(true)}>
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

      {getSections('action')}
      {getSections('reaction')}

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
                  onChange={event => setTitle(event.nativeEvent.text)}
                  keyboardType="default" // Adjust as needed
                />
              </View>
              <View style={[styles.optionRow]}>
                <Text style={[styles.optionLabel, { color: 'black' }]}>
                  Description
                </Text>
                <TextInput
                  style={styles.optionInput}
                  value={description}
                  onChange={event => setDescription(event.nativeEvent.text)}
                  keyboardType="default" // Adjust as needed
                />
              </View>
              <View style={[styles.optionRow]}>
                <Text style={[styles.optionLabel, { color: 'black' }]}>
                  Refresh rate
                </Text>
                <TextInput
                  style={[styles.optionInput]}
                  value={refreshRate ? String(refreshRate) : ''}
                  onChange={event =>
                    setRefreshRate(Number(event.nativeEvent.text))
                  }
                  keyboardType="numeric"
                />
              </View>
            </View>
            <View
              style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
              <TouchableOpacity onPress={handleSaveArea}>
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
                }}>
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

      {renderModal(
        isActionModalVisible,
        setIsActionModalVisible,
        selectedActionOptions,
        handleActionOptionChange,
        handleSaveAction,
        'Modify Action',
        area.action.service.color,
      )}

      {renderModal(
        isReactionModalVisible,
        setIsReactionModalVisible,
        selectedReactionOptions,
        handleReactionOptionChange,
        handleSaveReaction,
        'Modify Reaction',
        area.reaction.service.color,
      )}
      <BottomNavBar navigation={navigation} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    padding: 16,
  },
  subContainer: {
    flexDirection: 'column',
    marginBottom: 16,
    padding: 16,
    borderRadius: 10,
  },
  header: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 16,
    textAlign: 'center',
  },
  detailContainer: {
    flexDirection: 'row',
    marginBottom: 8,
  },
  ActionReactionHeader: {
    fontSize: 20,
    fontWeight: 'bold',
    textAlign: 'center',
    alignItems: 'center',
    marginBottom: 16,
  },
  label: {
    fontWeight: 'bold',
    fontSize: 16,
    color: 'white',
    marginRight: 8,
  },
  value: {
    color: 'white',
    fontSize: 16,
  },
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContent: {
    padding: 20,
    borderRadius: 10,
    width: '80%',
  },
  modalHeader: {
    fontSize: 20,
    fontWeight: 'bold',
    color: 'white',
    marginBottom: 16,
    textAlign: 'center',
  },
  input: {
    borderWidth: 1,
    borderColor: 'gray',
    borderRadius: 8,
    padding: 8,
    marginBottom: 16,
  },
  cancelButton: {
    borderWidth: 2,
    borderColor: '#E60000',
    backgroundColor: 'white',
    borderRadius: 8,
    padding: 4,
    alignItems: 'center',
    marginBottom: 8,
  },
  saveButton: {
    borderWidth: 2,
    borderColor: 'white',
    borderRadius: 8,
    padding: 4,
    alignItems: 'center',
  },
  optionRow: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 15,
    width: '100%',
  },
  optionLabel: {
    color: 'white',
    fontSize: 16,
    flex: 1,
  },
  optionInput: {
    flex: 2,
    backgroundColor: '#f0f0f0',
    borderRadius: 5,
    padding: 10,
    borderColor: '#ccc',
    borderWidth: 1,
  },
});

export default AreaDetailsScreen;
