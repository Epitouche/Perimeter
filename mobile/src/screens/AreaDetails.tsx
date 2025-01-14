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
import MdiPencilCircleOutline from '../components/icons/PencilCircleOutline';
import { AppContext } from '../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetailsScreen = ({ route }: Props) => {
  const { area } = route.params;
  const { ipAddress } = useContext(AppContext);
  const [isActionModalVisible, setIsActionModalVisible] = useState(false);
  const [isReactionModalVisible, setIsReactionModalVisible] = useState(false);
  const [selectedActionOptions, setSelectedActionOptions] = useState<{
    [key: string]: any;
  }>({});
  const [selectedReactionOptions, setSelectedReactionOptions] = useState<{
    [key: string]: any;
  }>({});

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

  const handleSaveAction = () => {
    fetch(`http://${ipAddress}:8080/api/v1/area`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        action_option: Object.values(selectedActionOptions),
        action_refresh_rate: area.action_refresh_rate,
        createdAt: area.createdAt,
        description: area.description,
        enable: area.enable,
        id: area.id,
        reaction_option: Object.values(selectedReactionOptions),
        storage_variable: area.storage_variable,
        title: area.title,
        update_at: new Date().toISOString(),
      }),
    })
      .then(response => response.json())
      .then(data => {
        console.log('Success:', data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
    setIsActionModalVisible(false);
  };

  const handleSaveReaction = () => {
    console.log(selectedReactionOptions);
    setIsReactionModalVisible(false);
  };

  React.useEffect(() => {
    const initialActionOptions = Object.entries(area.action_option).reduce(
      (acc, [name, value]) => {
        acc[name] = value;
        return acc;
      },
      {} as { [key: string]: any }
    );
    setSelectedActionOptions(initialActionOptions);
  }, [area.action_option]);

  for (const option of Object.entries(area.reaction_option).map(
    ([name, value]) => ({ name, value }),
  )) {
    selectedReactionOptions[option.name] = option.value;
  }

  console.log(selectedActionOptions);

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Area Details</Text>

      {/* Action Section */}
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
              <Text style={styles.label}>Action Service:</Text>
              <Text style={styles.value}>{area.action.service.name}</Text>
            </View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Action:</Text>
            </View>
          </View>
          <TouchableOpacity onPress={() => setIsActionModalVisible(true)}>
            <MdiPencilCircleOutline size={50} color="white" />
          </TouchableOpacity>
        </View>
      </View>

      {/* Reaction Section */}
      <View
        style={[
          styles.subContainer,
          { backgroundColor: area.reaction.service.color },
        ]}>
        <View style={styles.ActionReactionHeader}>
          <Text style={styles.label}>Reaction</Text>
        </View>
        <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
          <View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Reaction Service:</Text>
              <Text style={styles.value}>{area.reaction.service.name}</Text>
            </View>
            <View style={styles.detailContainer}>
              <Text style={styles.label}>Reaction:</Text>
            </View>
          </View>
          <TouchableOpacity onPress={() => setIsReactionModalVisible(true)}>
            <MdiPencilCircleOutline size={50} color="white" />
          </TouchableOpacity>
        </View>
      </View>

      {/* Action Modal */}
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
                  />
                </View>
              ))}
            </View>
            <View
              style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
              <TouchableOpacity onPress={handleSaveAction}>
                <View style={styles.saveButton}>
                  <Text style={[{ color: 'white' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity onPress={() => setIsActionModalVisible(false)}>
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

      {/* Reaction Modal */}
      <Modal
        visible={isReactionModalVisible}
        transparent={true}
        animationType="slide">
        <View style={styles.modalContainer}>
          <View
            style={[
              styles.modalContent,
              { backgroundColor: area.reaction.service.color },
            ]}>
            <Text style={styles.modalHeader}>Modify Reaction</Text>
            <View style={[{ flexDirection: 'column' }]}>
              {Object.keys(selectedReactionOptions).map(key => (
                <View key={key} style={styles.optionRow}>
                  <Text style={styles.optionLabel}>{key}</Text>
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
                  />
                </View>
              ))}
            </View>
            <View
              style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
              <TouchableOpacity onPress={handleSaveReaction}>
                <View style={styles.saveButton}>
                  <Text style={[{ color: 'white' }, { fontSize: 16 }]}>
                    Save
                  </Text>
                </View>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setIsReactionModalVisible(false)}>
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
