import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import BottomNavBar from '../../Components/NavBar';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../../App';

type Props = NativeStackScreenProps<RootStackParamList, 'WorkflowScreen'>;

const WorkflowScreen = ({navigation}: {navigation : any}) => {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Add Area</Text>
      <View style={styles.actionBox}>
        <Text style={styles.boxText}>Action</Text>
        <TouchableOpacity
        onPress={() => navigation.navigate('AddActionScreen')} 
        style={styles.addButton}>
          <Text style={styles.addText}>Add</Text>
        </TouchableOpacity>
      </View>
      <View style={styles.line} />
      <View style={styles.reactionBox}>
        <Text style={styles.boxText}>Reaction</Text>
        <TouchableOpacity style={styles.addButtonDisabled}>
          <Text style={styles.addTextDisabled}>Add</Text>
        </TouchableOpacity>
      </View>
      <BottomNavBar navigation={navigation} />
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
    backgroundColor: 'black',
    padding: 15,
    borderRadius: 8,
    width: '80%',
    marginBottom: 10,
  },
  reactionBox: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: 'gray',
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
    backgroundColor: '#ccc',
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

export default WorkflowScreen;
