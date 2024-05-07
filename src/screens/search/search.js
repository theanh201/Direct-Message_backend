import { View, Text, TouchableOpacity, TextInput, StyleSheet,FlatList, Image } from 'react-native'
import React from 'react'
import AntDesign from 'react-native-vector-icons/AntDesign'
import Colors from '../../asset/styles/color'
export default function SearchScreen() {
  const contactList = [
    {id:1, name: "Huy", img: require("../../asset/images/design/user.jpg")},
    {id:2, name: "Huy", img: require("../../asset/images/design/user.jpg")},
  ]
  const keywordHist = [
    {id:1, keyword: "Nhóm trọ vui vẻ"},
    {id:2, keyword: "Nhóm trọ vui vẻ"},
  ]
  return (
    <View>
      <View style={{flexDirection:"row",justifyContent:"space-between",backgroundColor:Colors._secondary, paddingHorizontal:15, paddingVertical:10, alignItems:"center"}}>
        <TouchableOpacity style={{}}>
          <AntDesign name="arrowleft" size={26} color={Colors._white}/>
        </TouchableOpacity>
        <TextInput style={{width:'90%',borderRadius:10,backgroundColor:Colors._white, paddingHorizontal:15}} placeholder="Tìm kiếm" />
      </View>
      <View style={styles.contact_list}>
        <Text>Liên hệ đã tìm</Text>
        <FlatList 
        data={contactList}
        horizontal={true}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item)=>item.id}
        renderItem={({item})=>(
          <View style={styles.contact_item}>
            <Image style={styles.contact_img} source={item.img}/>
            <Text style={styles.contact_name}>{item.name}</Text>
          </View>
        )}/>

      </View>
      <View style={styles.keyword_hist}>
        <Text style={{}}>Từ khóa đã tìm</Text>
        <FlatList 
        data={keywordHist}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item)=>item.id}
        renderItem={({item})=>(
          <View style={styles.keyword_item}>
            <AntDesign name="search1" color={Colors._black} size={20}/>
            <Text style={styles.keyword_txt}>{item.keyword}</Text>
          </View>
        )}/>
      </View>
    </View>
  )
}
const styles = StyleSheet.create({
  contact_list:{
    padding:15
  },
  contact_item:{
    alignItems:'center',
    marginHorizontal:5,
    marginTop:10
  },
  contact_img:{
    width:40,
    height:40,
    borderRadius:100,
    marginHorizontal:5
  },
  contact_name:{
    color:Colors._black
  },
  keyword_hist:{
    backgroundColor:Colors._pink,
    padding:15
  },
  keyword_item:{
    flexDirection:"row",
    alignItems:"center",
    marginVertical:5,
    backgroundColor:Colors._white,
    shadowColor:Colors._black,
    elevation:5,
    paddingHorizontal:15,
    paddingVertical:10,
    borderRadius:10
  },
  keyword_txt:{
    marginLeft:5
  }
})