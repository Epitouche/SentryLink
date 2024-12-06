import React from 'react';
import {View, Text, FlatList, TouchableOpacity, StyleSheet} from 'react-native';
import Svg, {Path} from 'react-native-svg';
import BottomNavBar from './NavBar';


const services = [
  {
    id: '1',
    name: 'Spotify',
    color: '#1DB954',
    icon: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M12 0a12 12 0 1012 12A12 12 0 0012 0zm5.33 17.46a.87.87 0 01-1.2.29 12.73 12.73 0 00-5.77-1.66 12.92 12.92 0 00-5.31 1.28.86.86 0 01-.82-1.53 14.68 14.68 0 016.13-1.46 14.41 14.41 0 016.53 1.86.87.87 0 01.29 1.22zm1.79-3.26a1.07 1.07 0 01-1.47.35 18.18 18.18 0 00-6.65-2.16 17.86 17.86 0 00-6.94 1.31 1.08 1.08 0 11-.84-2A20.77 20.77 0 0112 11a20.88 20.88 0 017.71 2.47 1.07 1.07 0 01.41 1.45zm.36-3.29a1.31 1.31 0 01-1.78.44A22.75 22.75 0 0012 9.66a22.92 22.92 0 00-8.68 1.7 1.31 1.31 0 01-1-.06 1.31 1.31 0 01.69-2.48A25.36 25.36 0 0112 7.34a25.37 25.37 0 019.62 2.06 1.32 1.32 0 01.44 1.86z"
        />
      </Svg>
    ),
  },
  {
    id: '2',
    name: 'Timer',
    color: '#9C27B0',
    icon: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M15.07 1H8.93V2.5H15.07V1ZM12 7A8 8 0 1 0 20 15 8 8 0 0 0 12 7ZM18.93 14H12V8.93A6.12 6.12 0 0 1 18.93 14Z"
        />
      </Svg>
    ),
  },
  {
    id: '3',
    name: 'Dropbox',
    color: '#2196F3',
    icon: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M5.64 3.6L10.73 6.7 5.64 9.8 0.55 6.7 5.64 3.6ZM18.36 3.6L23.45 6.7 18.36 9.8 13.27 6.7 18.36 3.6ZM0.55 12.2L5.64 15.3 10.73 12.2 5.64 9.1 0.55 12.2ZM18.36 9.1L23.45 12.2 18.36 15.3 13.27 12.2 18.36 9.1ZM12 12.2L16.64 15.3 12 18.4 7.36 15.3 12 12.2Z"
        />
      </Svg>
    ),
  },
  {
    id: '4',
    name: 'Weather Map',
    color: '#FF5722',
    icon: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M6 13A5 5 0 1 0 1 8H3A3 3 0 0 1 6 11 3 3 0 0 1 9 8H11A5 5 0 0 0 6 13ZM21 14H18.93A5 5 0 1 0 18.93 20H21V18H18.93A3 3 0 1 1 18.93 16H21V14Z"
        />
      </Svg>
    ),
  },
  {
    id: '5',
    name: 'Gmail',
    color: '#F44336',
    icon: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M12 13L3.5 6.5V18H20.5V6.5L12 13ZM3 4H21A1 1 0 0 1 22 5V19A1 1 0 0 1 21 20H3A1 1 0 0 1 2 19V5A1 1 0 0 1 3 4ZM20.5 5H3.5L12 10.5L20.5 5Z"
        />
      </Svg>
    ),
  },
];

const ServicesScreen = ({ navigation }: { navigation : any }) => {
  const renderService = ({ item }: { item: any }) => (
    <TouchableOpacity
      style={[styles.serviceButton, {backgroundColor: item.color}]}>
      {item.icon({width: 36, height: 36})}
      <Text style={styles.serviceText}>{item.name}</Text>
    </TouchableOpacity>
  );

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Services</Text>
      <View style={styles.separator} />
      <FlatList
        data={services}
        renderItem={renderService}
        keyExtractor={item => item.id}
        numColumns={2}
        contentContainerStyle={styles.list}
      />
      <BottomNavBar navigation={navigation} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    paddingHorizontal: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginTop: 20,
    textAlign: 'center',
  },
  separator: {
    height: 1,
    backgroundColor: '#ccc',
    marginVertical: 10,
  },
  list: {
    justifyContent: 'center',
    paddingVertical: 20,
  },
  serviceButton: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    margin: 10,
    height: 100,
    borderRadius: 10,
  },
  serviceText: {
    marginTop: 8,
    color: 'white',
    fontWeight: 'bold',
  },
});

export default ServicesScreen;
