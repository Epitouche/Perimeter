import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetailsScreen = ({ route }: Props) => {
  const { area } = route.params;

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Area Details</Text>
      <View style={styles.detailContainer}>
        <Text style={styles.label}>Action Service:</Text>
        <Text style={styles.value}>{area.action.service.name}</Text>
      </View>
      <View style={styles.detailContainer}>
        <Text style={styles.label}>Action:</Text>
        <Text style={styles.value}>{area.action.name}</Text>
      </View>
      <View style={styles.detailContainer}>
        <Text style={styles.label}>Reaction Service:</Text>
        <Text style={styles.value}>{area.reaction.service.name}</Text>
      </View>
      <View style={styles.detailContainer}>
        <Text style={styles.label}>Reaction:</Text>
        <Text style={styles.value}>{area.reaction.name}</Text>
      </View>
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
    textAlign: 'center',
  },
  detailContainer: {
    flexDirection: 'row',
    marginBottom: 8,
  },
  label: {
    fontWeight: 'bold',
    fontSize: 16,
    marginRight: 8,
  },
  value: {
    fontSize: 16,
  },
});

export default AreaDetailsScreen;
