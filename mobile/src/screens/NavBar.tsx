import React from 'react';
import { View, TouchableOpacity, StyleSheet } from 'react-native';
import { NavigationProp } from '@react-navigation/native';
import { SvgFromUri } from 'react-native-svg';

/**
 * BottomNavBar component renders a bottom navigation bar with four buttons.
 * Each button navigates to a different screen when pressed.
 *
 * @param {Object} props - The component props.
 * @param {NavigationProp<any>} props.navigation - The navigation prop used to navigate between screens.
 *
 * @returns {JSX.Element} The BottomNavBar component.
 */
const BottomNavBar = ({ navigation }: { navigation: NavigationProp<any> }) => {
  return (
    <View style={styles.navbarContainer}>
      <TouchableOpacity
        onPress={() => navigation.navigate('AreaView')}
        accessibilityLabel="Home navigation button"
        accessibilityHint="Navigates to the home screen"
        style={styles.navButton}>
        <SvgFromUri
          uri={'https://api.iconify.design/mdi:view-grid-outline.svg'}
          width={24}
          height={24}
          color={'#000000'}
        />
      </TouchableOpacity>
      <TouchableOpacity
        accessibilityLabel="Workflow navigation button"
        accessibilityHint="Navigates to the workflow screen"
        onPress={() => navigation.navigate('WorkflowScreen')}
        style={styles.navButton}>
        <SvgFromUri
          uri={'https://api.iconify.design/mdi:plus-box-outline.svg'}
          width={24}
          height={24}
          color={'#000000'}
        />
      </TouchableOpacity>
      <TouchableOpacity
        accessibilityLabel="Service navigation button"
        accessibilityHint="Navigates to the services screen"
        onPress={() => navigation.navigate('ServicesScreen')}
        style={styles.navButton}>
        <SvgFromUri
          uri={'https://api.iconify.design/mdi:login.svg'}
          width={24}
          height={24}
          color={'#000000'}
        />
      </TouchableOpacity>
      <TouchableOpacity
        accessibilityLabel="Settings navigation button"
        accessibilityHint="Navigates to the settings screen"
        onPress={() => navigation.navigate('SettingsScreen')}
        style={styles.navButton}>
        <SvgFromUri
          uri={'https://api.iconify.design/mdi:account.svg'}
          width={24}
          height={24}
          color={'#000000'}
        />
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  navbarContainer: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
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

export default BottomNavBar;
