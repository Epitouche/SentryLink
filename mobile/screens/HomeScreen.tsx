import React, {useState} from 'react';
import {
  View,
  Text,
  Button,
  TextInput,
  StyleSheet,
  StatusBar,
} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../App';

type Props = NativeStackScreenProps<RootStackParamList, 'Home'>;

const HomeScreen: React.FC<Props> = ({navigation}) => {
  const [ipAddress, setIpAddress] = useState<string>('');

  return (
    <View style={styles.container}>
      <Text>Enter the IP address to ping:</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter IP address"
        value={ipAddress}
        onChangeText={setIpAddress}
        keyboardType="numeric"
      />
      <Button title="Connect" onPress={() => navigation.navigate('Login', { 'ip': ipAddress })} />
    </View>
  );
};

export default HomeScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 16,
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    width: '100%',
    marginVertical: 10,
    paddingHorizontal: 10,
  },
  responseContainer: {
    marginTop: 20,
    padding: 10,
    backgroundColor: '#f0f0f0',
    borderRadius: 5,
    width: '100%',
  },
  responseText: {
    fontWeight: 'bold',
    marginBottom: 5,
  },
});
