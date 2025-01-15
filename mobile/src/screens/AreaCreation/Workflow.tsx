import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import BottomNavBar from '../NavBar';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'WorkflowScreen'>;

const WorkflowScreen = ({ navigation }: { navigation: any }) => {
  return (
    <View style={styles.container}>
      <Text
        style={styles.title}
        accessibilityLabel="Add Area Title"
        accessibilityHint="Displays the title of the screen">
        Add Area
      </Text>
      <View style={styles.actionBox}>
        <Text
          style={styles.boxText}
          accessibilityLabel="Action Text"
          accessibilityHint="Displays the action text">
          Action
        </Text>
        <TouchableOpacity
          onPress={() => navigation.navigate('AddActionScreen')}
          style={styles.addButton}
          accessibilityLabel="Add Action Button"
          accessibilityHint="Navigates to the Add Action screen">
          <Text style={styles.addText}>Add</Text>
        </TouchableOpacity>
      </View>
      <View style={styles.line} />
      <View style={styles.reactionBox}>
        <Text
          style={styles.boxText}
          accessibilityLabel="Reaction Text"
          accessibilityHint="Displays the reaction text">
          Reaction
        </Text>
        <TouchableOpacity
          style={styles.addButtonDisabled}
          accessibilityLabel="Add Reaction Button Disabled"
          accessibilityHint="Button is disabled and cannot be pressed">
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
