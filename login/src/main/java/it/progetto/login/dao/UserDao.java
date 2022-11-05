package it.progetto.login.dao;

import it.progetto.login.model.LoggedUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserDao extends JpaRepository<LoggedUser,Integer> {

}
