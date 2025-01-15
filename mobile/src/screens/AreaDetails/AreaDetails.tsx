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

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

const AreaDetailsScreen = ({ navigation, route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [isReactionModalVisible, setIsReactionModalVisible] = useState(false);
  const [selectedReactionOptions, setSelectedReactionOptions] = useState<{
    [key: string]: any;
  }>({});
  console.log('Area refreshrate:', area.action_refresh_rate);

  const [areaResults, setAreaResults] = useState([
    { created_at: '', result: '' },
  ]);

  const handleReactionOptionChange = (key: string, value: any, type: any) => {
    setSelectedReactionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
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

  for (const option of Object.entries(area.reaction_option).map(
    ([name, value]) => ({ name, value }),
  )) {
    selectedReactionOptions[option.name] = option.value;
  }

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
        {AreaSections({navigation, route})}

        {/* Action Section */}
        {ActionsSections({navigation, route})}

        {/* Reaction Section */}
        <View
          style={[
            styles.subContainer,
            { backgroundColor: area.reaction.service.color },
          ]}>
          <View style={styles.ActionReactionHeader}>
            <Text style={styles.label}>Reaction</Text>
          </View>
          <View
            style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
            <View>
              <View style={styles.detailContainer}>
                <Text style={styles.label}>Service:</Text>
                <Text style={styles.value}>{area.reaction.service.name}</Text>
              </View>
              <View style={styles.detailContainer}>
                <Text style={styles.label}>Options:</Text>
                <Text style={styles.value}>
                  {Object.entries(selectedReactionOptions).map(
                    ([key, value]) => `${key}: ${value} `,
                  )}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={() => setIsReactionModalVisible(true)}>
              <SvgFromUri
                uri={'https://api.iconify.design/mdi:pencil-circle-outline.svg'}
                width={50}
                height={50}
                color={'white'}
              />
            </TouchableOpacity>
          </View>
        </View>

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
                style={{
                  flexDirection: 'row',
                  justifyContent: 'space-between',
                }}>
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
                    <Text style={[{ color: '#E60000' }, { fontSize: 16 }]}>
                      Cancel
                    </Text>
                  </View>
                </TouchableOpacity>
              </View>
            </View>
          </View>
        </Modal>
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
