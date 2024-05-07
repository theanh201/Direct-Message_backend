import { View, Text, StyleSheet, TextInput, TouchableOpacity } from 'react-native'
import Icon from 'react-native-vector-icons/Ionicons';
import React from 'react'
import AntDesign from 'react-native-vector-icons/AntDesign'
import Colors from '../asset/styles/color';
export default function HeaderComponent({navigation}) {
  return (
    <View style={styles.header}>
    <AntDesign name='search1' size={24}/>
    <View style={styles.searchBar}>
      <TouchableOpacity
        style={styles.searchInput} onPress={()=>navigation.navigate("SearchScreen")}>
          <Text style={styles.txt_input}>Tìm kiếm</Text>
      </TouchableOpacity>
    </View>
    <Icon name="ellipsis-vertical" size={20} color="#333" style={styles.icon} />
  </View>
  )
}
const styles = StyleSheet.create({
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 15,
    paddingVertical:10,
    borderBlockColor:Colors._secondary,
    borderBottomWidth:1
  },
  searchBar: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#f0f0f0',
    borderRadius: 10,
    marginHorizontal:10,
    paddingHorizontal:5
  },
  searchInput: {
    flex: 1,
    fontSize: 16,
    color: '#333',
    paddingHorizontal:15,
    paddingVertical:10
  },
  txt_input:{
    fontSize:14,
    fontWeight:'500',
  }
})