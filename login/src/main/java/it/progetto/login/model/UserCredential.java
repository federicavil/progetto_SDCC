package it.progetto.login.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import lombok.NonNull;

import javax.persistence.*;

@Entity
@Data
@JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
public class UserCredential {

    @Id
    private String username;
    @NonNull
    private String password;
    @OneToOne(cascade = CascadeType.ALL)
    private LoggedUser user;

    public UserCredential(){

    }

    public UserCredential(String username, String password, LoggedUser user){
        this.password = password;
        this.user = user;
        this.username = username;
    }

    public String getPassword() {
        return this.password;
    }

    public LoggedUser getUser() {
        return this.user;
    }
}
