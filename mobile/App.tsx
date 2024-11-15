import { StatusBar } from 'expo-status-bar';
import { View, Text, Button, StyleSheet } from 'react-native';
import { useState } from 'react';

export default function App() {
  const [responseData, setResponseData] = useState<string | null>(null);

  const handleButtonClick = async () => {
    try {
      const response = await fetch('http://10.134.197.212:8080/ping');
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
      <Text>Open up App.tsx to start working on your app!</Text>
      <Button title="Click me" onPress={handleButtonClick} />
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
