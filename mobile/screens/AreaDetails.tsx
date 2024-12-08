import React from 'react';
import {View, Text, StyleSheet, ScrollView, TouchableOpacity} from 'react-native';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../App';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetails = ({navigation, route}: Props) => {
  const {area} = route.params;

  return (
    <View>
      <ScrollView style={styles.container}>
        <Text style={styles.header}>Area Details</Text>

        <View style={styles.detailSection}>
          <Text style={styles.detailTitle}>Action</Text>
          <View style={styles.detailContent}>
            <MaterialCommunityIcons
              name={getIconName(area.action_id.service_id.name)}
              size={24}
              color="black"
              style={styles.icon}
            />
            <View>
              <Text style={styles.detailText}>Name: {area.action_id.name}</Text>
              <Text style={styles.detailText}>
                Description: {area.action_id.description}
              </Text>
              <Text style={styles.detailText}>
                Service: {area.action_id.service_id.name}
              </Text>
            </View>
          </View>
        </View>

        <View style={styles.detailSection}>
          <Text style={styles.detailTitle}>Reaction</Text>
          <View style={styles.detailContent}>
            <MaterialCommunityIcons
              name={getIconName(area.reaction_id.service_id.name)}
              size={24}
              color="black"
              style={styles.icon}
            />
            <View>
              <Text style={styles.detailText}>
                Name: {area.reaction_id.name}
              </Text>
              <Text style={styles.detailText}>
                Description: {area.reaction_id.description}
              </Text>
              <Text style={styles.detailText}>
                Service: {area.reaction_id.service_id.name}
              </Text>
            </View>
          </View>
        </View>

        <View style={styles.detailSection}>
          <Text style={styles.detailTitle}>Additional Information</Text>
          <Text style={styles.detailText}>
            Enabled: {area.enable ? 'Yes' : 'No'}
          </Text>
          <Text style={styles.detailText}>
            Created At: {new Date(area.createdAt).toLocaleString()}
          </Text>
          <Text style={styles.detailText}>
            Updated At: {new Date(area.update_at).toLocaleString()}
          </Text>
        </View>
      </ScrollView>
      <TouchableOpacity
        style={styles.backButton}
        onPress={() => navigation.goBack()}>
        <Text style={styles.backButtonText}>Back</Text>
      </TouchableOpacity>
    </View>
  );
};

const getIconName = (serviceName: string) => {
  switch (serviceName.toLowerCase()) {
    case 'timer':
      return 'clock-outline';
    case 'spotify':
      return 'spotify';
    case 'dropbox':
      return 'dropbox';
    case 'gmail':
      return 'gmail';
    default:
      return 'help-circle-outline';
  }
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
  detailSection: {
    marginBottom: 24,
  },
  detailTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 8,
  },
  detailContent: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  icon: {
    marginRight: 12,
  },
  detailText: {
    fontSize: 16,
    marginBottom: 4,
  },
  backButton: {
    marginTop: 20,
    width: '90%',
    height: 50,
    borderRadius: 25,
    borderWidth: 2,
    borderColor: '#000',
    justifyContent: 'center',
    alignItems: 'center',
  },
  backButtonText: {
    fontSize: 18,
    fontWeight: 'bold',
  },
});

export default AreaDetails;
