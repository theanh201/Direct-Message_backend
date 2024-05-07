import React, { useState } from 'react';
import { View, TextInput, Button, StyleSheet, Alert } from 'react-native';

const LoginScreen = ({navigation}) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const goToSignup = () => {
    navigation.navigate('SignupScreen');
  };
  const handleLogin = () => {
    // Basic email and password validation
    if (!validateEmail(email)) {
        Alert.alert('Email không chính xác');
        return;
      }
    if (!email.trim() || !password.trim()) {
      Alert.alert("Xin vui lòng nhập đầy đủ Email và Mật Khẩu");
      return;
    }

    // Add your authentication logic here
    console.log('Email:', email);
    console.log('Password:', password);
    navigation.navigate('HomeScreen');
  };
  const validateEmail = (email) => {
    // Email validation regex pattern
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };
  return (
    <View style={styles.container}>
        <Button title="Go to Signup" onPress={goToSignup} />
      <TextInput
        style={styles.input}
        placeholder="Email"
        keyboardType="email-address"
        onChangeText={text => setEmail(text)}
        value={email}
        autoCapitalize="none"
      />
      <TextInput
        style={styles.input}
        placeholder="Password"
        secureTextEntry
        onChangeText={text => setPassword(text)}
        value={password}
      />
      <Button title="Login" onPress={handleLogin} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 20,
  },
  input: {
    height: 40,
    width: '100%',
    borderColor: 'gray',
    borderWidth: 1,
    borderRadius: 5,
    marginBottom: 10,
    paddingHorizontal: 10,
  },
});

export default LoginScreen;
