import React, { useState, useEffect } from 'react';
import { View, Text, ActivityIndicator, StyleSheet, ScrollView, Image, TouchableOpacity } from 'react-native';
import axios from 'axios';
import Colors from '../../asset/styles/color';
import Feather from 'react-native-vector-icons/Feather';
const MyComponent = () => {
  const [students, setStudents] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get('http://demo-api.stecom.vn:8888/api/student/get-all?pageSize=20&pageIndex=1');

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
              <Image style={styles.img} source={require('../../asset/images/design/user.jpg')}/>
              <View style={styles.info}>
                <Text style={styles.name}>{student.name}</Text>
                <Text style={styles.dob}>{student.dateOfBirth}</Text>
              </View>
              <View style={{flexDirection:'row', justifyContent:'space-between', alignItems:'center', width:'20%'}}>
                <TouchableOpacity>
                  <Feather name='phone-call' size={20} color={Colors._green}/>
                </TouchableOpacity>
                <TouchableOpacity>
                  <Feather name='message-circle' size={20} color={Colors._secondary}/>
                </TouchableOpacity>
              </View>
            </View>
          ))
        ) : (
          <Text>No students available</Text>
        )}
      </View>
      <View style={{width:'100%', height:100}}>

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
    justifyContent:'space-between',
    alignItems:'center',
    width:'100%',
    padding:15
  },
  studentContainer: {
    width:'100%',
    flexDirection:'row',
    justifyContent:'space-between',
    alignItems:'center',
    marginBottom: 10,
    padding: 10,
    borderRadius: 5,
    backgroundColor:Colors._white,
    shadowColor:Colors._black,
    elevation:2
  },
  img:{
    width:60,
    height:60,
    borderRadius:10
  },
  info:{
    marginLeft: 10
  },
  name: {
    fontSize: 14,
    fontWeight: 'bold',
    color: Colors._secondary,
  },
  dob:{
    fontSize: 12,
    color: Colors._black
  }
});

export default MyComponent;
