package com.shenali.user.service.repository;

import com.shenali.user.service.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {

    User findByUserId(Long userId);

    User[] findByDepartmentId(Long departmentId);
}
