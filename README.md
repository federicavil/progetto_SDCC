# Progetto SDCC: Mattia Antonangeli, Federica Villani

Per effettuare il deploy del progetto via Docker Compose:


0) Eseguire "git clone https://github.com/federicavil/progetto_SDCC.git" per scaricare il progetto

1) Modificare la variabile HOST_HOME_ADDRESS nel file ".env" presente nella root del progetto con il path della directory in cui è presente 
    la cartella ".aws" relativa alle credenziali di accesso ad AWS;
    
2) Effettuare la build di Docker Compose dalla root del progetto:

        $ docker compose build
        
3) Eseguire il deploy del progetto tramite Docker Compose:

        $ docker-compose up

Per modificare le porte esposte dai container si possono modificare i seguenti file nella sezione [docker_compose] relativa:

- Frontend            --> conf/host.ini
- API_Gateway         --> conf/host.ini
- NotificationManager --> conf/conf.ini
- PathManager         --> conf/database.ini
- WeatherForecast     --> conf/conf.ini
- VisitManager        --> conf/database.ini
