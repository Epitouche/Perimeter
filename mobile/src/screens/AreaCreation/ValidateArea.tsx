import React, { useContext, useEffect } from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  StyleSheet,
  TextInput,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'ValidateAreaScreen'>;

const ValidateAreaScreen: React.FC<Props> = ({ navigation, route }) => {
  const { actionId, actionOptions, reactionId, reactionOptions } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  interface Service {
    color: string;
    created_at: string;
    description: string;
    icon: string;
    id: number;
    name: string;
    oauth: boolean;
    update_at: string;
  }

  const [actionService, setActionService] = React.useState<Service | null>(
    null,
  );
  const [reactionService, setReactionService] = React.useState<Service | null>(
    null,
  );
  const [actionName, setActionName] = React.useState('');
  const [reactionName, setReactionName] = React.useState('');
  const [title, setTitle] = React.useState('');
  const [description, setDescription] = React.useState('');
  const [refreshTimer, setRefreshTimer] = React.useState('');

  useEffect(() => {
    const getService = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/action/info/${actionId}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const res = await fetch(
          `http://${ipAddress}:8080/api/v1/reaction/info/${reactionId}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          },
        );

        let actionData = await response.json();
        let reactionData = await res.json();
        console.log('reactionData', reactionData);
        setActionName(actionData[0].name);
        setReactionName(reactionData[0].name);
        setActionService(actionData[0].service);
        setReactionService(reactionData[0].service);
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching service:', error);
      }
    };
    getService();
  }, [token, ipAddress, actionId, reactionId]);

  const saveButtonPressed = async () => {
    try {
      console.log(parseInt(refreshTimer));
      let data = await fetch(`http://${ipAddress}:8080/api/v1/area`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          action_id: actionId,
          action_option: actionOptions,
          action_refresh_rate: parseInt(refreshTimer),
          description,
          reaction_id: reactionId,
          reaction_option: reactionOptions,
          title,
        }),
      });
      let res = await data.json();
      navigation.navigate('AreaView');
    } catch (error) {
      if (error.code === 401) {
        navigation.navigate('Login');
      }
      console.error('Error saving area:', error);
    }
  };

  return (
    <View style={styles.container}>
      <View style={{ width: '80%', marginBottom: 20 }}>
      <Text style={styles.label}>Title</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter title"
        onChangeText={text => setTitle(text)}
        accessibilityLabel="Title Input"
      />
      </View>
      <View style={{ width: '80%', marginBottom: 20 }}>
      <Text style={styles.label}>Description</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter description"
        onChangeText={text => setDescription(text)}
        accessibilityLabel="Description Input"
      />
      </View>
      <View style={{ width: '80%', marginBottom: 20 }}>
      <Text style={styles.label}>Refresh Timer (in seconds)</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter refresh timer"
        keyboardType="numeric"
        onChangeText={text => setRefreshTimer(text)}
        accessibilityLabel="Refresh Timer Input"
      />
      </View>
      <Text style={styles.title}>Add Area</Text>
      <View
      style={[
        styles.actionBox,
        { backgroundColor: actionService?.color, borderRadius: 8 },
      ]}>
      <Text style={styles.boxText} accessibilityLabel="Action Name">
        {actionName}
      </Text>
      </View>
      <View style={styles.line} />
      <View
      style={[
        styles.actionBox,
        { backgroundColor: reactionService?.color, borderRadius: 8 },
      ]}>
      <Text style={styles.boxText} accessibilityLabel="Reaction Name">
        {reactionName}
      </Text>
      </View>
      <TouchableOpacity
      style={styles.saveButton}
      onPress={() => {
        saveButtonPressed();
      }}
      accessibilityLabel="Save Button"
      >
      <Text style={styles.addText}>Save</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#fff',
  },
  title: {
    fontSize: 24,
    marginBottom: 20,
  },
  actionBox: {
    flexDirection: 'row',
    backgroundColor: 'black',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 15,
    borderRadius: 8,
    width: '80%',
    marginBottom: 10,
  },
  boxText: {
    color: '#fff',
    fontSize: 18,
  },
  addButton: {
    backgroundColor: 'white',
    paddingVertical: 5,
    paddingHorizontal: 15,
    borderRadius: 5,
  },
  addButtonDisabled: {
    paddingVertical: 5,
    paddingHorizontal: 15,
    borderRadius: 5,
  },
  addText: {
    color: 'black',
    fontSize: 16,
    fontWeight: 'bold',
  },
  addTextDisabled: {
    color: 'gray',
    fontSize: 16,
    fontWeight: 'bold',
  },
  line: {
    width: 2,
    height: 20,
    backgroundColor: 'black',
    marginVertical: 10,
  },
  saveButton: {
    backgroundColor: '#4CAF50',
    paddingVertical: 10,
    paddingHorizontal: 20,
    borderRadius: 8,
    marginTop: 20,
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    width: '100%',
    marginVertical: 5,
    paddingHorizontal: 10,
    color: 'black',
  },
  label: {
    fontSize: 16,
    marginBottom: 2,
  },
});

export default ValidateAreaScreen;
