package com.shenali.user.service.service;

import com.shenali.user.service.entity.User;
import com.shenali.user.service.repository.UserRepository;
import com.shenali.user.service.valueobjects.Department;
import com.shenali.user.service.valueobjects.ResponseTemplateVO;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.ArrayList;

@Service
@Slf4j
public class UserService {
    @Autowired
    private UserRepository userRepository;

    @Autowired
    private RestTemplate restTemplate;

    public User saveUser(User user) {
        log.info("saveUser : UserService");
        return userRepository.save(user);
    }

    public ResponseTemplateVO getUserWithDepartment(Long userId) {
        log.info("getUserWithDepartment : UserService");

        ResponseTemplateVO vo = new ResponseTemplateVO();

        User user = userRepository.findByUserId(userId);
        Department department = restTemplate.getForObject("http://DEPARTMENT-SERVICE/departments/" + user.getDepartmentId(), Department.class);

        vo.setUser(user);
        vo.setDepartment(department);

        return vo;
    }

//    public User[] getUsersByDepartmentId(Long departmentId) {
//        log.info("getUsersByDepartmentId : UserService");
//        return userRepository.findByDepartmentId(departmentId);
//    }

    public ArrayList<ResponseTemplateVO> getUsersByDepartmentId(Long departmentId) {
        log.info("getUsersByDepartmentId : UserService");

        ArrayList<ResponseTemplateVO> responses = new ArrayList<>();
        User[] users = userRepository.findByDepartmentId(departmentId);

        for (User user: users) {
            ResponseTemplateVO response = new ResponseTemplateVO();
            response.setUser(user);
            Department department = restTemplate.getForObject("http://DEPARTMENT-SERVICE/departments/" + user.getDepartmentId(), Department.class);
            response.setDepartment(department);
            responses.add(response);
        }

        return responses;
    }
}
