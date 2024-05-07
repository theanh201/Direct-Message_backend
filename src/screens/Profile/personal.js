import { View, Text, ImageBackground, StyleSheet, Image, TouchableOpacity, FlatList, SafeAreaView, ScrollView } from 'react-native'
import React, { startTransition,useRef, useState} from 'react'
import Colors from '../../asset/styles/color'
import Entypo from 'react-native-vector-icons/Entypo'
import { launchImageLibrary } from 'react-native-image-picker';
import { launchCamera } from 'react-native-image-picker';


export default function PersonalScreen({navigation}) {
    const [expanded, setExpanded] = useState(false);
    const [numLines, setNumLines] = useState(0);
  
    const toggleExpand = () => {
        setExpanded(!expanded);
      };

    const liststory = [
        {id:1, date:"28/04/2024", describe:"Cập nhật ảnh đại diện", img:require('../../asset/images/design/avt.jpg')},
        {id:2, date:"09/11/2023", describe:"Hải Phònglà một trong năm thành phố trực thuộc trung ương của Việt Nam. Đây là thành phố lớn thứ ba Việt Nam, thành phố cảng quan trọng, trung tâm công nghiệp, cảng biển, đồng thời cũng là trung tâm kinh tế, văn hóa, y tế, giáo dục, khoa học, thương mại và công nghệ thuộc Vùng duyên hải Bắc Bộ của Việt Nam.", img:require('../../asset/images/design/user.jpg')},
        {id:3, date:"09/11/2023", describe:"hello world", img:require('../../asset/images/design/user.jpg')},
        {id:4, date:"09/11/2023", describe:"hello world", img:require('../../asset/images/design/user.jpg')},
        {id:5, date:"09/11/2023", describe:"hello world", img:require('../../asset/images/design/bg.jpg')},
        {id:6, date:"09/11/2023", describe:"hello world", img:require('../../asset/images/design/user.jpg')},
        {id:7, date:"09/11/2023", describe:"hello world", img:require('../../asset/images/design/user.jpg')}
    ]
    const openImagePicker = () => {
        const options = {
          mediaType: 'photo',
          includeBase64: false,
          maxHeight: 2000,
          maxWidth: 2000,
        };
    
        launchImageLibrary(options, (response) => {
          if (response.didCancel) {
            console.log('User cancelled image picker');
          } else if (response.error) {
            console.log('Image picker error: ', response.error);
          } else {
            let imageUri = response.uri || response.assets?.[0]?.uri;
            setSelectedImage(imageUri);
          }
        });
      };

      handleCameraLaunch = () => {
        const options = {
          mediaType: 'photo',
          includeBase64: false,
          maxHeight: 2000,
          maxWidth: 2000,
        };
      
        launchCamera(options, response => {
          if (response.didCancel) {
            console.log('User cancelled camera');
          } else if (response.error) {
            console.log('Camera Error: ', response.error);
          } else {
            let imageUri = response.uri || response.assets?.[0]?.uri;
            setSelectedImage(imageUri);
            console.log(imageUri);
          }
        });
      }
    
  return (
    <>
        <ImageBackground source={require('../../asset/images/design/bg.jpg')} style={styles.personal_bg}>
            <TouchableOpacity style={styles.personal_update_avatar } onPress={handleCameraLaunch}>
                <Image source={require('../../asset/images/design/avt.jpg')} style={styles.personal_avatar}/>
            </TouchableOpacity>
            <View style={{justifyContent:'space-between', flexDirection:'row', paddingHorizontal:15, width:'100%'}}>
                <TouchableOpacity style={styles.btn} onPress={()=>navigation.goBack()}>
                    <Entypo name='chevron-left' color={Colors._white} size={24}/>      
                </TouchableOpacity>
                <TouchableOpacity style={styles.btn} onPress={()=>navigation.navigate("SettingScreen")}>
                    <Entypo name='dots-three-horizontal' color={Colors._white} size={24}/>      
                </TouchableOpacity>
            </View>
            </ImageBackground>
        <View style={{paddingTop:80, paddingHorizontal:15, width:'100%', alignItems:'center'}}>
            <TouchableOpacity style={{justifyContent:'space-between', flexDirection:'row', alignItems:'center'}} onPress={()=>{navigation.navigate('Txt')}}>
                <Entypo name='pencil' size={16}/>
                <Text style={styles.personal_text}>Cập nhật giới thiệu về bản thân</Text>
            </TouchableOpacity>
            <View style={{backgroundColor:Colors._white, width:'100%',padding:20, margin:20, justifyContent:'space-between', flexDirection:'row',borderRadius:10}}>
                <Text>Viết nhật ký</Text>
                <Entypo name="folder-images" size={20} color={Colors._pink}/>
            </View>
        </View>
        <View style={{alignItems:'center', width:'100%'}}>
            <FlatList
            data={liststory}
            scrollEnabled={true}
            keyExtractor={(item)=>item.id}
            renderItem={({item})=>(
                <View style={styles.story_item}>
                    <Text style={styles.txt_date}>{item.date}</Text>
                    <Text 
                    numberOfLines={expanded?undefined:5} 
                    style={styles.txt_content}
                    onLayout={(event) => {
                        const { height } = event.nativeEvent.layout;
                        const lineHeight = StyleSheet.flatten(styles.txt_content).lineHeight;
                        const calculatedNumLines = Math.floor(height / lineHeight);
                        setNumLines(calculatedNumLines);}}>
                        {item.describe}
                    </Text>
                    {!expanded && numLines === 5 &&(
                        <TouchableOpacity onPress={toggleExpand}>
                            <Text style={styles.showMore}>Xem thêm</Text>
                        </TouchableOpacity>
                    )}
                    <Image source={item.img} style={styles.story_img}/>
                </View>
            )}/>

        </View>
        <View style={{height:500, width:'100%'}}>
            <Text>lkajfkdsf</Text>
        </View>
    </>
  )
}
const styles = StyleSheet.create({
    personal_bg:{
        width:'100%',
        height:200,
        position:'relative'
    },
    personal_update_avatar:{
        position:'absolute',
        bottom:-70,
        left:'32%',
    },
    personal_avatar:{
        width:140,
        height:140,
        borderRadius:100,
        
        borderWidth:4,
        borderColor:Colors._white
    },
    personal_text:{
        fontSize:14,
        fontWeight:'500'
    },
    btn:{
        width:50,
        height:50,
        alignItems:'center',
        justifyContent:'center',
        borderRadius:50,
        
    },
    story_item:{
        width:350,
        height:'auto',
        backgroundColor:Colors._white,
        padding:15,
        paddingVertical:10,
        borderRadius:10,
        marginVertical:10,
    },
    txt_date:{
        fontSize:10
    },
    txt_content:{
        marginVertical:10,
        lineHeight:20,
        color:Colors._black
    },
    story_img:{
        marginTop:10,
        width:160,
        height:160,
        borderRadius:10
    },
    showMore: {
        color: 'blue',
        marginTop: 5,
    },
})