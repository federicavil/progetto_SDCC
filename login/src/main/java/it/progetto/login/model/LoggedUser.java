package it.progetto.login.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
@Data
@JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
public class LoggedUser {
    @Id
    @GeneratedValue
    private Integer id;

    private String name;
    private String surname;
    private String description;

    public LoggedUser(){

    }

    public LoggedUser(String name, String surname, String description){
        this.description = description;
        this.name = name;
        this.surname = surname;
    }

    public Integer getId(){
        return this.id;
    }

}
