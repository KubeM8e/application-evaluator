package com.shenali.user.service.controller;

import com.shenali.user.service.entity.User;
import com.shenali.user.service.service.UserService;
import com.shenali.user.service.valueobjects.ResponseTemplateVO;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;

@RestController
@RequestMapping("/users")
@Slf4j
public class UserController {
    @Autowired
    private UserService userService;

    @PostMapping("/")
    public User saveUser(@RequestBody User user) {
        log.info("saveUser : UserController");
        return userService.saveUser(user);
    }

    @GetMapping("/{id}")
    public ResponseTemplateVO getUserWithDepartment(@PathVariable("id") Long userId) {
        log.info("getUserWithDepartment : UserController");
        return userService.getUserWithDepartment(userId);
    }

//    @GetMapping("")
//    public User[] getUsersByDepartmentId(@RequestParam("depId") Long departmentId) {
//        log.info("getUserByDepartmentId : UserController");
//        return userService.getUsersByDepartmentId(departmentId);
//    }

    @GetMapping("")
    public ArrayList<ResponseTemplateVO> getUsersByDepartmentId(@RequestParam("depId") Long departmentId) {
        log.info("getUserByDepartmentId : UserController");
        return userService.getUsersByDepartmentId(departmentId);
    }


}
