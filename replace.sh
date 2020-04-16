
#!/bin/bash
files=( 
"gen/simple/servicehandler.go"
"gen/simple/serviceinterface.go"
)

replacements=( "s/\/\/jsonplaceholder/\/jsonplaceholder/" )

for f in "${files[@]}"; do
        for d in "${replacements[@]}"; do
        sed -i "" "$d" $f
        done
done

        # "s/type varchar (50)/type ENUM('this', 'that')/"
# 
# sed -i "" "s/time varchar (50)/time DATETIME/" spec.sql


# sed -i "" "s/name varchar (50)/name char(100)/" spec.sql
# # sed -i "" "s/type varchar (50)/type ENUM('this', 'that')/" spec.sql
# sed -i "" "s/location_type varchar (50)/location_type ENUM('wing', 'entrance')/" spec.sql
# sed -i "" "s/payment_type varchar (50)/payment_type ENUM('credit', 'debit', 'cash')/" spec.sql

# sed -i "" "s/audio_type varchar (50)/audio_type ENUM('physical', 'app')/" spec.sql



# sed -i "" "s/arrival_transport varchar (50)/arrival_transport ENUM('Walk', 'Metro', 'Train', 'Bus', 'Taxi', 'Private Vehicle', 'Other')/" spec.sql
# # ('Pyramid', 'Carrousel de Louvre', '99 rue de Rivoli', 'Passage Richelieu','Portes de Lion')
# sed -i "" "s/listen_time varchar (50)/listen_time TIME/" spec.sql
# sed -i "" "s/is_online varchar (50)/is_online BOOLEAN/" spec.sql
# sed -i "" "s/price char (100)/price decimal/" spec.sql
# sed -i "" "s/varchar (50)/char (100)/" spec.sql

# sed -i "" "s/language varchar (50)/language ENUM ('French', 'English', 'Italian', 'Russian', 'Chinese' 'Mandarin', 'Spanish', 'Portuguese', 'Korean', 'Dutch', 'Polish', 'Swedish', 'Norwegian', 'Finnish')/" spec.sql



