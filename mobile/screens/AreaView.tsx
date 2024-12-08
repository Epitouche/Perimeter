import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import BottomNavBar from './NavBar';
import {RootStackParamList} from '../App';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaView'>;

const AreasScreen = ({ navigation, route }: Props) => {
  const areas = [
    { text: 'Upload every day', color: '#FF4D4D', icons: ['github', 'clock-outline'] },
    { text: 'Start Music!', color: '#4CAF50', icons: ['cloud-outline', 'spotify'] },
    { text: 'Upload every day', color: '#9C27B0', icons: ['cloud-upload-outline', 'clock-outline'] },
    { text: 'Stock photo!', color: '#2196F3', icons: ['gmail', 'dropbox'] },
  ];

  return (
    <View style={styles.container}>
      <Text style={styles.header}>My AREAs</Text>
      <View style={styles.areasContainer}>
        {areas.map((area, index) => (
          <View key={index} style={[styles.areaBox, { backgroundColor: area.color }]}>
            <Text style={styles.areaText}>{area.text}</Text>
            <View style={styles.iconsContainer}>
              {area.icons.map((icon, idx) => (
                <MaterialCommunityIcons key={idx} name={icon} size={24} color="white" style={styles.areaIcon} />
              ))}
            </View>
          </View>
        ))}
      </View>
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
  header: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 16,
  },
  areasContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-around',
  },
  areaBox: {
    width: 150,
    height: 150,
    borderRadius: 16,
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 16,
  },
  areaText: {
    color: 'white',
    fontSize: 16,
    fontWeight: 'bold',
    textAlign: 'center',
    marginBottom: 8,
  },
  iconsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    width: '80%',
  },
  areaIcon: {
    marginHorizontal: 4,
  },
  navbarContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
    paddingVertical: 8,
    backgroundColor: '#f0f0f0',
    borderTopWidth: 1,
    borderTopColor: '#d0d0d0',
  },
  navButton: {
    alignItems: 'center',
    justifyContent: 'center',
  },
});

export default AreasScreen;
