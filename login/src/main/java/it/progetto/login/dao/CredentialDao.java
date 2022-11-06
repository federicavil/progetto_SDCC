package it.progetto.login.dao;

import it.progetto.login.model.UserCredential;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CredentialDao extends JpaRepository<UserCredential,String> {

    UserCredential findByUsername(String username);

}
