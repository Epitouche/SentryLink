import { StatusBar } from 'expo-status-bar';
import { View, Text, Button, TextInput, StyleSheet } from 'react-native';
import { useState } from 'react';

export default function App() {
  const [responseData, setResponseData] = useState<string | null>(null);
  const [ipAddress, setIpAddress] = useState<string>(''); // State to hold the IP address

  const handleButtonClick = async () => {
    try {
      if (!ipAddress) {
        setResponseData('Please enter a valid IP address');
        return;
      }

      const response = await fetch(`http://${ipAddress}:8080/ping`);
      const data = await response.json();
      console.log('Response:', data);
      setResponseData(JSON.stringify(data, null, 2)); // Set the response data to state
    } catch (error) {
      console.error('Error making GET request:', error);
      setResponseData('Error fetching data');
    }
  };

  return (
    <View style={styles.container}>
      <Text>Enter the IP address to ping:</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter IP address"
        value={ipAddress}
        onChangeText={setIpAddress} // Update the IP address state
        keyboardType="numeric" // Restrict keyboard to numeric inputs
      />
      <Button title="Ping IP" onPress={handleButtonClick} />
      {responseData && (
        <View style={styles.responseContainer}>
          <Text style={styles.responseText}>Response:</Text>
          <Text>{responseData}</Text>
        </View>
      )}
      <StatusBar style="auto" />
    </View>
  );
}

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
