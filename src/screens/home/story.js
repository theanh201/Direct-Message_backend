import React, { useState } from 'react';
import { StyleSheet, View, Text, ScrollView, TouchableOpacity, Image } from 'react-native';
import Colors from '../../asset/styles/color';

const StoryScreen = () => {
  const [stories, setStories] = useState([
    { id: 1, user_name: 'BKISS', user_avt:require('../../asset/images/design/bkiss.jpg'), unread: true, txt_content:'🌟Exploring English Vocabulary for Historical Holidays!🌟 \
    🌞Chào mừng các bạn đã đến với câu lạc bộ tiếng anh BEC. Hôm nay chúng ta hãy cùng nhau khám phá về chủ đề từ vựng tiếng anh cho ngày nghỉ lễ liên quan đến lịch sử. \
    💁Những ngày nghỉ lễ không chỉ là thời gian nghỉ ngơi mà còn là cơ hội để chúng ta kỷ niệm và tôn vinh những sự kiện, nhân vật quan trọng trong lịch sử. \
    ❤️Chúng ta thật may mắn khi được sinh ra trong thời kỳ hòa bình, vậy nên hãy mang tấm', img_content: require('../../asset/images/design/bg.jpg')},
    // Add more stories as needed
    { id: 2, user_name: 'BKISS', user_avt:require('../../asset/images/design/bkiss.jpg'), unread: true, txt_content:'🌟Exploring English Vocabulary for Historical Holidays!🌟 \
    🌞Chào mừng các bạn đã đến với câu lạc bộ tiếng anh BEC. Hôm nay chúng ta hãy cùng nhau khám phá về chủ đề từ vựng tiếng anh cho ngày nghỉ lễ liên quan đến lịch sử. \
    💁Những ngày nghỉ lễ không chỉ là thời gian nghỉ ngơi mà còn là cơ hội để chúng ta kỷ niệm và tôn vinh những sự kiện, nhân vật quan trọng trong lịch sử. \
    ❤️Chúng ta thật may mắn khi được sinh ra trong thời kỳ hòa bình, vậy nên hãy mang tấm', img_content: require('../../asset/images/design/bg.jpg')},
  ]);

  const markAsRead = (id) => {
    const updatedStories = stories.map(story =>
      story.id === id ? { ...story, unread: false } : story
    );
    setStories(updatedStories);
  };

  return (
    <View style={styles.container}>
      <ScrollView showsHorizontalScrollIndicator={false} contentContainerStyle={styles.storiesContainer}>
        {stories.map(story => (
          <View key={story.id} style={[styles.story_item, { borderColor: story.unread ? '#3b5998' : 'transparent' }]}>
            <View style={{flexDirection:'row', alignItems:'center'}}>
              <Image source={story.user_avt} style={styles.story_uavt} />
              <Text style={styles.story_uname}>{story.user_name}</Text>
            </View>
            <Text style={styles.story_txt}>
              {story.txt_content}
            </Text>
            <Image style={styles.story_img} source={story.img_content}/>

          </View>
        ))}
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  story_item:{
    padding:15,
    backgroundColor:Colors._white
  },
  story_uavt:{
    width:50,
    height:50,
    borderRadius:60
  },
  story_uname:{
    color:Colors._black,
    fontSize:14,
    fontWeight:'bold',
    marginLeft:10
  },
  story_img:{
    width:'100%',
    height:250,
    borderRadius:10
  }
});

export default StoryScreen;
