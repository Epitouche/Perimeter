import React, { useContext } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'ValidateAreaScreen'>;

const ValidateAreaScreen: React.FC<Props> = ({ navigation, route }) => {
  const { actionId, actionOptions, reactionId, reactionOptions } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  let service: any;
  let reactionService: any;
  const [actionName, setActionName] = React.useState('');
  const [reactionName, setReactionName] = React.useState('');

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

      const serviceResponse = await fetch(
        `http://${ipAddress}:8080/api/v1/action/info/service/${actionId}`,
        {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        },
      );
      const serviceRes = await fetch(
        `http://${ipAddress}:8080/api/v1/reaction/info/service/${reactionId}`,
        {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        },
      );

      setActionName((await response.json())[0].name);
      setReactionName((await res.json())[0].name);
      service = (await serviceResponse.json())[0];
      reactionService = (await serviceRes.json())[0];

    } catch (error) {
      if (error.code === 401) {
        navigation.navigate('Login');
      }
      console.error('Error fetching service:', error);
    }
  };
  getService();

  const saveButtonPressed = async () => {
    try {
      let data = await fetch(`http://${ipAddress}:8080/api/v1/area`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          action_id: actionId,
          actionOptions,
          reaction_id: reactionId,
          reactionOptions,
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
      <Text style={styles.title}>Add Area</Text>
      <View
        style={[
          styles.reactionBox,
          { backgroundColor: service?.color || '#000' },
        ]}>
        <Text style={styles.boxText}>{actionName}</Text>
      </View>
      <View style={styles.line} />
      <View
        style={[
          styles.reactionBox,
          { backgroundColor: reactionService?.color || '#000' },
        ]}>
        <Text style={styles.boxText}>{reactionName}</Text>
      </View>
      <TouchableOpacity
        style={styles.addButton}
        onPress={() => {
            saveButtonPressed();
        }}>
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
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 15,
    borderRadius: 8,
    width: '80%',
    marginBottom: 10,
  },
  reactionBox: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 15,
    borderRadius: 8,
    width: '80%',
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
});

export default ValidateAreaScreen;
