import React, { useState, useEffect } from 'react';
import { View, Text, ActivityIndicator, StyleSheet, ScrollView } from 'react-native';
import axios from 'axios';

const MyComponent = () => {
  const [students, setStudents] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get('http://demo-api.stecom.vn:8888/api/student/get-all?pageSize=10&pageIndex=1');

      const extractedData = response.data.items.map(item => ({
        id: item.id,
        name: item.name,
        dateOfBirth: item.dateOfBirth,
      }));
      setStudents(extractedData);
      setLoading(false);
    } catch (error) {
      console.error('Error fetching data:', error);
      setLoading(false);
    }
  };

  return (
    <ScrollView contentContainerStyle={styles.scrollViewContainer}>
      <View style={styles.container}>
        {loading ? (
          <ActivityIndicator size="large" color="#0000ff" />
        ) : students.length > 0 ? (
          students.map(student => (
            <View key={student.id} style={styles.studentContainer}>
              <Text style={styles.id}>ID: {student.id}</Text>
              <Text style={styles.name}>Name: {student.name}</Text>
              <Text style={styles.score}>Score: {student.dateOfBirth}</Text>
            </View>
          ))
        ) : (
          <Text>No students available</Text>
        )}
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  scrollViewContainer: {
    flexGrow: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  container: {
    flexDirection:'row',
    justifyContent:'space-between',
    alignItems:'center',
    flexWrap:'wrap',
    width:'100%',
    padding:15
  },
  studentContainer: {
    width:'45%',
    marginBottom: 10,
    borderWidth: 1,
    padding: 10,
    borderRadius: 5,
  },
  id: {
    fontSize: 16,
    fontWeight: 'bold',
    color: 'blue',
  },
  name: {
    fontSize: 16,
    fontWeight: 'bold',
    color: 'green',
  },
  score: {
    fontSize: 16,
    fontWeight: 'bold',
    color: 'red',
  },
});

export default MyComponent;
