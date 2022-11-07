package it.progetto.login.model;

import com.fasterxml.jackson.annotation.JsonAlias;
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
    @JsonAlias("Name")
    private String name;
    @JsonAlias("Surname")
    private String surname;
    @JsonAlias("Description")
    private String description;

    public LoggedUser(){

    }

    public LoggedUser(String name, String surname, String description){
        this.description = description;
        this.name = name;
        this.surname = surname;
    }

    public LoggedUser(Integer id, String name, String surname, String description){
        this.description = description;
        this.name = name;
        this.surname = surname;
        this.id = id;
    }

    public Integer getId(){
        return this.id;
    }
    public String getName(){return this.name;}

    @Override
    public String toString(){
        return "\"name\":\""+ name + "\",\"surname\":\"" + surname + "\",\"description\":\""+description+"\"";
    }

}
