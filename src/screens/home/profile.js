import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native'
import React from 'react'
import Colors from '../../asset/styles/color'
import SwitchUser from 'react-native-vector-icons/MaterialCommunityIcons'
import Entypo from 'react-native-vector-icons/Entypo'
import PersonalScreen from '../Profile/personal'
import { createNativeStackNavigator } from '@react-navigation/native-stack'

export default function UserScreen({navigation}) {
  return (
    <View>
      <View style={{flexDirection:'row', justifyContent:'space-between', padding:15, alignItems:'center', backgroundColor:Colors._secondary}}>
        <TouchableOpacity style={{flexDirection:'row'}} onPress={()=>navigation.navigate("PersonalScreen")}>
          <Image style={{width:50, height:50, borderRadius:20}} source={require('../../asset/images/design/avt.jpg')}/>
          <View style={{marginHorizontal:10}}>
            <Text style={{color:Colors._white, fontWeight:'bold'}}>Đào Đức Huy</Text>
            <Text style={{color:Colors._dash, fontSize:12}}>Xem trang cá nhân</Text>
          </View>
        </TouchableOpacity>
        <TouchableOpacity>
          <SwitchUser name="account-switch" size={25} color={Colors._white}/>
        </TouchableOpacity>
      </View>

      <TouchableOpacity style={styles.profile_box}>
        <View style={styles.profile_content}>
          <Entypo name="lock" size={24} color={Colors._secondary}/>
          <Text style={styles.profile_text}>Tài khoản và bảo mật</Text>
        </View>
        <Entypo name="chevron-right" size={24} color={Colors._secondary}/>
      </TouchableOpacity>
      <TouchableOpacity style={styles.profile_box}>
        <View style={styles.profile_content}>

          <Entypo name="shield" size={24} color={Colors._secondary}/>
          <Text style={styles.profile_text}>Tài khoản và bảo mật</Text>
        </View>
        <Entypo name="chevron-right" size={24} color={Colors._secondary}/>
      </TouchableOpacity>
    </View>
  )
}
const styles = StyleSheet.create({
  profile_box:{
    flexDirection:'row',
    justifyContent:'space-between',
    alignItems:'center',
    backgroundColor:Colors._white,
    width:"100%",
    padding:15,
    borderBottomWidth:1,
    borderColor: Colors._dash
  },
  profile_content:{
    flexDirection:'row',
  },
  profile_text:{
    fontWeight:'bold',
    fontSize:14,
    marginHorizontal:10
  }
})