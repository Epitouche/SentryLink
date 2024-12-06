import React from 'react';
import {View, TouchableOpacity, StyleSheet} from 'react-native';
import Svg, {Path} from 'react-native-svg';
import {NavigationProp} from '@react-navigation/native';

const BottomNavBar = ({navigation}: {navigation: NavigationProp<any>}) => {
  return (
    <View style={styles.navbarContainer}>
      <TouchableOpacity
        onPress={() => navigation.navigate('AreaView')}
        style={styles.navButton}>
        <Svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24">
          <Path
            fill="currentColor"
            d="M4 8h4V4H4zm6 12h4v-4h-4zm-6 0h4v-4H4zm0-6h4v-4H4zm6 0h4v-4h-4zm6-10v4h4V4zm-6 4h4V4h-4zm6 6h4v-4h-4zm0 6h4v-4h-4z"
          />
        </Svg>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => navigation.navigate('Add')}
        style={styles.navButton}>
        <Svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24">
          <Path
            fill="currentColor"
            d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2m0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8s8 3.59 8 8s-3.59 8-8 8"
          />
        </Svg>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => navigation.navigate('ServiceScreen')}
        style={styles.navButton}>
        <Svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24">
          <Path
            fill="currentColor"
            d="M11 7L9.6 8.4l2.6 2.6H2v2h10.2l-2.6 2.6L11 17l5-5zm9 12h-8v2h8c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2h-8v2h8z"
          />
        </Svg>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => navigation.navigate('Profile')}
        style={styles.navButton}>
        <Svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24">
          <Path
            fill="currentColor"
            d="M12 5.9a2.1 2.1 0 1 1 0 4.2a2.1 2.1 0 0 1 0-4.2m0 9c2.97 0 6.1 1.46 6.1 2.1v1.1H5.9V17c0-.64 3.13-2.1 6.1-2.1M12 4C9.79 4 8 5.79 8 8s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4m0 9c-2.67 0-8 1.34-8 4v3h16v-3c0-2.66-5.33-4-8-4"
          />
        </Svg>
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