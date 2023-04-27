package com.shenali.department.service.controller;

import com.shenali.department.service.entity.Department;
import com.shenali.department.service.entity.User;
import com.shenali.department.service.service.DepartmentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/departments")
@Slf4j
public class DepartmentController {
    @Autowired
    private DepartmentService departmentService;

    // controller => service => repository

    @PostMapping("/")
    public Department saveDepartment(@RequestBody Department department) {
        log.info("savaDepartment : DepartmentController");
        return departmentService.savaDepartment(department);
    }

    @GetMapping("/{id}") // path variable
    public Department findDepartmentById (@PathVariable("id") Long departmentId) {
        log.info("findDepartment : DepartmentController");
        return departmentService.findDepartmentById(departmentId);

    }
}
