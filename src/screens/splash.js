import { View, Text, Image, StyleSheet } from 'react-native'
import React from 'react'
import Colors from '../asset/styles/color'

export default function SplashScreen({navigation}) {
    setTimeout(()=>{
        navigation.navigate("LoginScreen")
    },3000)

    return (
    <View style={{alignItems:'center', padding:14}}>
      <Image style={{width:340, height:200, resizeMode:'cover', borderRadius:20}} source={require('../asset/images/design/splash.jpeg')}/>
      <Text style={styles.splash_text}>cHAAT</Text>
      <View style={{flexDirection:'row', width:'80%', justifyContent:'space-between'}}>
        <View style={styles.splash_box}>
            <Text style={styles.splash_primary_txt}>c</Text>
            <Text style={styles.splash_text}>hat</Text>
        </View>
        <View style={styles.splash_box}>
            <Text style={styles.splash_primary_txt}>H</Text>
            <Text style={styles.splash_text}>uy</Text>
        </View>
        <View style={styles.splash_box}>
            <Text style={styles.splash_primary_txt}>A</Text>
            <Text style={styles.splash_text}>n</Text>
        </View>
        <View style={styles.splash_box}>
            <Text style={styles.splash_primary_txt}>A</Text>
            <Text style={styles.splash_text}>nh</Text>
        </View>
        <View style={styles.splash_box}>
            <Text style={styles.splash_primary_txt}>T</Text>
            <Text style={styles.splash_text}>iến</Text>
        </View>
      </View>
    </View>
    )
}
const styles = StyleSheet.create({
    splash_text:{
        color:'black',
        fontWeight:'bold',
        fontSize:20
    },
    splash_box:{
        flexDirection:'row',
        alignItems:'baseline'
    },
    splash_primary_txt:{
        fontSize:30,
        fontWeight:'bold',
        color:Colors._primary
    }
})