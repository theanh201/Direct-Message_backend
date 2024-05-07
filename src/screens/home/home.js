import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { StatusBar, View, Text, TextInput, TouchableOpacity, FlatList, StyleSheet, Image } from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';
import AntDesign from 'react-native-vector-icons/AntDesign';
import Entypo from 'react-native-vector-icons/Entypo';
import Colors from './../../asset/styles/color';
import MessageScreen from './message';
import StoryScreen from './story';
import CallScreen from './call';
import Profile from './profile';
import Friends from './friends';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HeaderComponent from '../../components/header';

const HomeScreen = ({navigation}) => {
  const theme = 'light';
  const Tab = createBottomTabNavigator();
  const [selectedTab, setSelectedTab] = React.useState('message');
  const [optionVisible, setOptionVisible] = React.useState(false);

  const toggleOptionVisible = () => {
    setOptionVisible(!optionVisible)
  }

  const navigateToScreen = (screen) => {
    setSelectedTab(screen);
    navigation.navigate(screen);
  };
  const screenOptions = {
    tabBarShowLabel:false,
    headerShown:false,
    tabBarStyle:{
      position: "absolute",
      bottom: 0,
      right: 0,
      left: 0,
      elevation: 0,
      height: 60,
      background: "#fff"
    }
  }

  return (
    
    <View style={styles.container}>
      <StatusBar backgroundColor={theme === "light" ? Colors._secondary : Colors._black} barStyle={theme === "light" ? "dark-content" : "light-content"}/>
      {/* HEADER */}
      <View style={styles.header}>
        <AntDesign name='search1' size={24}/>
        <View style={styles.searchBar}>
          <TouchableOpacity
            style={styles.searchInput} onPress={()=>navigation.navigate("SearchScreen")}>
              <Text style={styles.txt_input}>Tìm kiếm</Text>
          </TouchableOpacity>
        </View>
        <TouchableOpacity onPress={toggleOptionVisible}>
          <Entypo name="plus" size={24} color="#333" style={styles.icon} />
        </TouchableOpacity>
        
      </View>
      {optionVisible && <View style={styles.header_options}>
          <TouchableOpacity style={styles.option_fn}>
            <AntDesign name="adduser"/>
            <Text>Thêm bạn</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.option_fn}>
            <AntDesign name="addusergroup"/>
            <Text>Tạo nhóm</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.option_fn}>
            <AntDesign name="delete"/>
            <Text>Xóa hội thoại</Text>
          </TouchableOpacity>
        </View>}
      <Tab.Navigator screenOptions={screenOptions}>
        <Tab.Screen name="Messages" 
        component={MessageScreen} 
        options=
        {{
          tabBarIcon: ({focused})=>
          {
            return (
              <View style={styles.bottombar_item}> 
                <Entypo name="message" size={20} color={focused ? Colors._secondary: "#333"} />
                <Text style={[styles.bottombar_text, focused && styles.bottombar_focused]}>Nhắn tin</Text>
              </View>)
          }
        }}/>

        <Tab.Screen name="Friends" 
        component={Friends} 
        options=
        {{
          tabBarIcon: ({focused})=>
          {
            return (
              <View style={styles.bottombar_item}> 
                <Entypo name="users" size={20} color={focused ? Colors._secondary: "#333"}/>
                <Text style={[styles.bottombar_text, focused && styles.bottombar_focused]}>Danh bạ</Text>
              </View>)
          }
        }}/>
        <Tab.Screen name="Story" component={StoryScreen}
        options={{
          tabBarIcon: ({focused})=>
          {
            return (
              <View style={styles.bottombar_item}>
                <AntDesign name="star"  size={20} color={focused ? Colors._secondary: "#333"}/>
                <Text style={[styles.bottombar_text, focused && styles.bottombar_focused]}>Nhật ký</Text>
              </View>
            )
          }
        }}/>
        <Tab.Screen name="Profile" component={Profile}
        options={{
          tabBarIcon: ({focused})=>
          {
            return (
              <View style={styles.bottombar_item}>
                <Entypo name="user"  size={20} color={focused ? Colors._secondary: "#333"}/>
                <Text style={[styles.bottombar_text, focused && styles.bottombar_focused]}>Cá nhân</Text>
              </View>
            )
          }
        }}/>
      </Tab.Navigator>

      
    </View>
  );
};

const styles = StyleSheet.create({
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 15,
    paddingVertical:10,
    borderBlockColor:Colors._secondary,
    borderBottomWidth:1,
  },
  header_options:{
    position:'absolute',
    top:0,
    left:0,
    backgroundColor:Colors._secondary,
    padding:5,
    height:600,
    justifyContent:'space-between',
  },
  option_fn:{
    flexDirection:'row',
    justifyContent:'space-between',
    alignItems:'center'
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
  },
  container: {
    flex: 1,
    backgroundColor: '#fff',
    position:'relative'
  },
  profileButton: {
    marginRight: 10,
  },
  profileButtonText: {
    fontSize: 16,
    color: Colors._primary,
  },
 
  icon: {
    marginLeft: 10,
  },
  chatItem: {
    flexDirection:'row',
    alignItems:'center',
    paddingHorizontal: 20,
    paddingVertical: 15,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },

  navigation: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
    paddingVertical: 15,
    borderTopWidth: 1,
    borderTopColor: '#eee',
  },
  navigationItem: {
    flex: 1,
    alignItems: 'center',
  },
  navigationText: {
    fontSize: 16,
    color: Colors._primary,
    fontWeight:'bold'
  },
  selectedNavigationItem: {
    color:Colors._primary
  },
  selectedNavigationText: {
    fontWeight: 'bold',
  },
  home_bottombar:{

  },
  bottombar_item:{
    alignItems:'center'
  },
  bottombar_text:{
    fontSize:14,
    color:Colors._black,
    fontWeight:'bold'
  },
  bottombar_focused: {
    color:Colors._secondary // color when focused
  },
});

export default HomeScreen;
