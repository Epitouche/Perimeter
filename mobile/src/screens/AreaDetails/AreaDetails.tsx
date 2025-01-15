import React, { useState, useContext, useEffect } from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  Modal,
  TextInput,
  FlatList,
  ScrollView,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';
import { SvgFromUri } from 'react-native-svg';
import { styles } from './StylesAreaDetails';
import AreaSections from './AreaSections';
import ActionsSections from './ActionsSections';
import ReactionsSections from './ReactionsSections';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetailsScreen = ({ navigation, route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [areaResults, setAreaResults] = useState([
    { created_at: '', result: '' },
  ]);

  useEffect(() => {
    const fetchAreaResults = async () => {
      console.log('Fetching area results');
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/area-result/${area.id}`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        if (response.ok) {
          const body = await response.json();
          setAreaResults(body);
          console.log('Area results:', body);
        }
      } catch (error) {
        console.error('Error fetching area results:', error);
      }
    };
    fetchAreaResults();
  }, []);

  const renderItem = ({
    item,
  }: {
    item: { created_at: string; result: string };
  }) => (
    <View style={styles.row}>
      <Text style={styles.cell}>{item.created_at.substring(0, 19)}</Text>
      <Text style={styles.cell}>{item.result}</Text>
    </View>
  );

  return (
    <ScrollView>
      <View style={styles.container}>
        <Text style={styles.header}>Area Details</Text>

        {/* Area Section */}
        {AreaSections({ navigation, route })}

        {/* Action Section */}
        {ActionsSections({ navigation, route })}

        {/* Reaction Section */}
        {ReactionsSections({ navigation, route })}

        {/* Area Results */}
        <Text style={[styles.header, { marginTop: 16 }]}>Area Results</Text>
        <View
          style={[
            { borderColor: 'black', borderWidth: 1, borderRadius: 10, flex: 1 },
          ]}>
          <View style={[styles.row, { backgroundColor: 'white' }]}>
            <Text style={styles.cell}>created_at</Text>
            <Text style={styles.cell}>result</Text>
          </View>
          <FlatList
            data={areaResults}
            renderItem={renderItem}
            keyExtractor={item => item.created_at}
            style={{ flex: 1 }}
          />
        </View>

        <TouchableOpacity onPress={() => navigation.navigate('AreaView')}>
          <View style={{ alignItems: 'flex-end', justifyContent: 'flex-end' }}>
            <Text
              style={[
                styles.cancelButton,
                { color: '#E60000', width: '20%', margin: 10 },
              ]}>
              Back
            </Text>
          </View>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
};

export default AreaDetailsScreen;
