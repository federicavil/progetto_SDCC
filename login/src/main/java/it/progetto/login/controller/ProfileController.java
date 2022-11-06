package it.progetto.login.controller;

import com.google.gson.Gson;
import it.progetto.login.dao.UserDao;
import it.progetto.login.model.LoggedUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProfileController {

    @Autowired
    private UserDao userDao;

    public ProfileController(){}

    public String getUserProfile(Integer userId){
        LoggedUser user = userDao.findLoggedUserById(userId);
        return "{"+user.toString()+"}";
    }

    public String updateUserProfile(Integer userId, String userProfile){
        userProfile = "{\"id\": " + userId +","+userProfile.substring(1);
        System.out.println(userProfile);
        Gson gson = new Gson();
        LoggedUser user = gson.fromJson(userProfile, LoggedUser.class);
        System.out.println(user);
        System.out.println(user.getId());
        LoggedUser updateUser = userDao.save(user);
        return "{"+updateUser.toString()+"}";
    }
}
