package com.shenali.department.service.service;

import com.shenali.department.service.entity.Department;
import com.shenali.department.service.entity.User;
import com.shenali.department.service.repository.DepartmentRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
@Slf4j
public class DepartmentService {
    @Autowired
    private DepartmentRepository departmentRepository;

    public Department savaDepartment(Department department) {
        log.info("saveDepartment : DepartmentService");
        return departmentRepository.save(department);
    }

    public Department findDepartmentById(Long departmentId) {
        log.info("saveDepartment : DepartmentService");
        return departmentRepository.findDepartmentById(departmentId);
    }
}
