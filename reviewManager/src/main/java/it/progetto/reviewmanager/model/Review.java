package it.progetto.reviewmanager.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
@Data
@JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
public class Review {

    @Id
    @GeneratedValue
    public Long id;

    public String title;
    public String mountainPathName;
    public String comment;
    public Integer vote;
    public String author;

    @Override
    public String toString() {
        return "\"title\":\"" + title + "\",\"mountainPathName\":\"" + mountainPathName + "\",\"comment\":\"" + comment + "\"" + "\",\"vote\":\"" + vote + "\",\"author\":\"" + author + "\"";

    }
}
