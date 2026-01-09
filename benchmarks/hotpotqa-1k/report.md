# RagTune Retrieval Report

**Timestamp:** 2025-12-23T16:00:34Z  
**Collection:** hotpotqa-ollama  
**Store:** qdrant  

---

## Summary

| Config | Top-K | Recall@K | MRR | Coverage | Redundancy |
|--------|-------|----------|-----|----------|------------|
| top3 | 3 | 0.983 | 0.998 | 0.987 | 1.53 |
| top5 | 5 | 0.990 | 0.998 | 0.997 | 2.52 |
| top10 | 10 | 1.000 | 0.998 | 1.000 | 5.03 |
| top20 | 20 | 1.000 | 0.998 | 1.000 | 10.05 |

---

## Detailed Results

### Config: top3 (top_k=3)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt, Eski_Imaret_Mosque_b4d62c.txt | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt | ✅ |
| q2 | Random_House_Tower_2b6a2a.txt, 888_7th_Avenue_f091f0.txt, 888_7th_Avenue_f091f0.txt | 888_7th_Avenue_f091f0.txt, Random_House_Tower_2b6a2a.txt | ✅ |
| q3 | Alex_Ferguson_6962e9.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt, Peter_Schmeichel_0b144a.txt | 1995_96_Manchester_United_F.C._season_3f9756.txt, Alex_Ferguson_6962e9.txt | ✅ |
| q4 | Apple_Remote_fcaffa.txt, Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt | Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt | ✅ |
| q5 | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt, Charles_Nungesser_67ba12.txt | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt | ✅ |
| q6 | Henry_J._Kaiser_53c448.txt, Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt | Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt | ✅ |
| q7 | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt, Jean-Loup_Chrétien_1d1a07.txt | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt | ✅ |
| q8 | Freakonomics__film__249214.txt, In_the_Realm_of_the_Hackers_367f43.txt, Connections__TV_series__deb419.txt | In_the_Realm_of_the_Hackers_367f43.txt, Freakonomics__film__249214.txt | ✅ |
| q9 | Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt | Socialist_Revolutionary_Party_68737a.txt, Russian_Civil_War_5a6752.txt | ✅ |
| q10 | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt, Ogallala__Nebraska_3578d6.txt | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt | ✅ |
| q11 | Giuseppe_Arimondi_0fc7eb.txt, Battle_of_Adwa_1fa890.txt, Addis_Ababa_6c835e.txt | Battle_of_Adwa_1fa890.txt, Giuseppe_Arimondi_0fc7eb.txt | ✅ |
| q12 | Dirleton_Castle_a042b8.txt, Yellowcraigs_0c745c.txt, Kingdom_of_Northumbria_4036b8.txt | Yellowcraigs_0c745c.txt, Dirleton_Castle_a042b8.txt | ✅ |
| q13 | English_Electric_Canberra_becfbe.txt, Avro_Vulcan_7aa981.txt, English_Electric_Canberra_becfbe.txt | English_Electric_Canberra_becfbe.txt, No._2_Squadron_RAAF_a75482.txt | ✅ |
| q14 | Euromarché_085339.txt, Carrefour_a2a75b.txt, Maxeda_f76dfe.txt | Euromarché_085339.txt, Carrefour_a2a75b.txt | ✅ |
| q15 | Delirium__Ellie_Goulding_album__5bb0cb.txt, On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Erika_Jayne_60e347.txt | On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Delirium__Ellie_Goulding_album__5bb0cb.txt | ✅ |
| q16 | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt, The_Legend_of_Korra_86cdc2.txt | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt | ✅ |
| q17 | Oranjegekte_8248a5.txt, Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt | Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt | ✅ |
| q18 | Tromeo_and_Juliet_86fab3.txt, James_Gunn_9a06f8.txt, Romeo_87e7c1.txt | James_Gunn_9a06f8.txt, Tromeo_and_Juliet_86fab3.txt | ✅ |
| q19 | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt, Bob_Seger_69d05d.txt | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt | ✅ |
| q20 | Rostker_v._Goldberg_b61238.txt, Conscription_in_the_United_States_e920fa.txt, Conscription_in_the_United_States_e920fa.txt | Conscription_in_the_United_States_e920fa.txt, Rostker_v._Goldberg_b61238.txt | ✅ |
| q21 | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt, Orange_Julius_e106c8.txt | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt | ✅ |
| q22 | Their_Lives_eed122.txt, Monica_Lewinsky_7bb6c2.txt, Nancy_Soderberg_a37451.txt | Monica_Lewinsky_7bb6c2.txt, Their_Lives_eed122.txt | ✅ |
| q23 | Teide_National_Park_aaf674.txt, Garajonay_National_Park_97f362.txt, Hatton_Castle__Angus_eb96ea.txt | Garajonay_National_Park_97f362.txt, Teide_National_Park_aaf674.txt | ✅ |
| q24 | Andrew_Jaspan_f6dc15.txt, Andrew_Jaspan_f6dc15.txt, The_Conversation__website__724191.txt | The_Conversation__website__724191.txt, Andrew_Jaspan_f6dc15.txt | ✅ |
| q25 | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt, The_Conversation__website__724191.txt | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt | ✅ |
| q26 | Tysons_Galleria_ead975.txt, Oldham_County__Kentucky_34ed53.txt, Ogallala__Nebraska_3578d6.txt | Tysons_Galleria_ead975.txt, McLean__Virginia_45ec68.txt | ✅ |
| q27 | My_Eyes__Blake_Shelton_song__fa840b.txt, Based_on_a_True_Story..._42790b.txt, The_Best_of_LeAnn_Rimes_722f10.txt | Based_on_a_True_Story..._42790b.txt, My_Eyes__Blake_Shelton_song__fa840b.txt | ✅ |
| q28 | Caroline_Carver__actress__7577ed.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Carol__film__cbfb1f.txt | The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Caroline_Carver__actress__7577ed.txt | ✅ |
| q29 | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt, Wilton_Mall_66971e.txt | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt | ✅ |
| q30 | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt, Jessica_Rothe_59deb4.txt | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt | ✅ |
| q31 | Mummulgum_1292e7.txt, Casino__New_South_Wales_8c85c5.txt, Ogallala__Nebraska_3578d6.txt | Casino__New_South_Wales_8c85c5.txt, Mummulgum_1292e7.txt | ✅ |
| q32 | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt, Sacred_Planet_26fd7b.txt | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt | ✅ |
| q33 | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt, James_Gunn_9a06f8.txt | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt | ✅ |
| q34 | Roberta_Vinci_f714c8.txt, Jorge_Lozano_1baa7a.txt, Jorge_Lozano_1baa7a.txt | Jorge_Lozano_1baa7a.txt, Roberta_Vinci_f714c8.txt | ✅ |
| q35 | Marco_Da_Silva__dancer__777e91.txt, Erika_Jayne_60e347.txt, Cressida_Bonas_5c3c57.txt | Erika_Jayne_60e347.txt, Marco_Da_Silva__dancer__777e91.txt | ✅ |
| q36 | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt | ✅ |
| q37 | Kingdom_of_the_Isles_9a036f.txt, Kingdom_of_the_Isles_9a036f.txt, Aonghus_Mór_b5e643.txt | Aonghus_Mór_b5e643.txt, Kingdom_of_the_Isles_9a036f.txt | ✅ |
| q38 | Bruce_Spizer_c78d7e.txt, Bob_Seger_69d05d.txt, The_Beatles_c9e770.txt | The_Beatles_c9e770.txt, Bruce_Spizer_c78d7e.txt | ✅ |
| q39 | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt, Wayne_Garland_02f688.txt | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt | ✅ |
| q40 | Argand_lamp_563fb5.txt, Lewis_lamp_ddcc57.txt, Lewis_lamp_ddcc57.txt | Lewis_lamp_ddcc57.txt, Argand_lamp_563fb5.txt | ✅ |
| q41 | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt, Kathy_Sullivan__Australian_politician__a2272b.txt | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt | ✅ |
| q42 | Bishop_Carroll_Catholic_High_School_5aa312.txt, Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt | Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt | ✅ |
| q43 | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt, United_States_presidential_election__2016_363dfd.txt | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt | ✅ |
| q44 | Southaven__Mississippi_7cc376.txt, Memphis_Hustle_71bf0b.txt, Ogallala__Nebraska_3578d6.txt | Memphis_Hustle_71bf0b.txt, Southaven__Mississippi_7cc376.txt | ✅ |
| q45 | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt, Something_There_380707.txt | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt | ✅ |
| q46 | Albertina_eec665.txt, Hanna_Varis_e45643.txt, Hanna_Varis_e45643.txt | Albertina_eec665.txt, Hanna_Varis_e45643.txt | ✅ |
| q47 | Hatton_Hill_027377.txt, Hatton_Castle__Angus_eb96ea.txt, Hatton_Castle__Angus_eb96ea.txt | Hatton_Castle__Angus_eb96ea.txt, Hatton_Hill_027377.txt | ✅ |
| q48 | The_Legend_of_Korra_86cdc2.txt, Kuvira_4fdaa0.txt, Gargoyles__TV_series__acb424.txt | Kuvira_4fdaa0.txt, The_Legend_of_Korra_86cdc2.txt | ✅ |
| q49 | The_Five_Obstructions_18f018.txt, The_Importance_of_Being_Icelandic_086a7b.txt, The_Importance_of_Being_Icelandic_086a7b.txt | The_Importance_of_Being_Icelandic_086a7b.txt, The_Five_Obstructions_18f018.txt | ✅ |
| q50 | Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt, Will__amp__Grace_31b102.txt, The_Legend_of_Korra_86cdc2.txt | Will__amp__Grace_31b102.txt, Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt | ✅ |
| q51 | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt, Kohlberg_Kravis_Roberts_728df7.txt | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt | ✅ |
| q52 | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt, Wendell_Berry_4313ab.txt | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt | ✅ |
| q53 | 1920__film__db5e98.txt, Soha_Ali_Khan_b0280e.txt, 1920__film_series__78d752.txt | 1920__film_series__78d752.txt, 1920__film__db5e98.txt | ✅ |
| q54 | 71st_Golden_Globe_Awards_8dbed6.txt, 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt | 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt | ✅ |
| q55 | Charles_Hastings_Judd_409cf1.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt, Kalākaua_055843.txt | Charles_Hastings_Judd_409cf1.txt, Kalākaua_055843.txt | ✅ |
| q56 | Armie_Hammer_7c1778.txt, The_Polar_Bears_b06086.txt, Eddie_Izzard_fb2be0.txt | The_Polar_Bears_b06086.txt, Armie_Hammer_7c1778.txt | ✅ |
| q57 | 712_Fifth_Avenue_4807f9.txt, Manhattan_Life_Insurance_Building_fb6468.txt, Random_House_Tower_2b6a2a.txt | Manhattan_Life_Insurance_Building_fb6468.txt, 712_Fifth_Avenue_4807f9.txt | ✅ |
| q58 | Tenerife_c81266.txt, Gerald_Reive_c5cd23.txt, Samoa_086113.txt | Samoa_086113.txt, Gerald_Reive_c5cd23.txt | ✅ |
| q59 | Tecumseh_bcf68b.txt, Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt | Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt | ✅ |
| q60 | Samuel_Sim_07d71f.txt, Tromeo_and_Juliet_86fab3.txt, Bedknobs_and_Broomsticks_090f32.txt | Samuel_Sim_07d71f.txt, Awake__film__360ee6.txt | ✅ |
| q61 | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt, Øresund_Region_549ba5.txt | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt | ✅ |
| q62 | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt, Larry_Drake_85028c.txt | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt | ✅ |
| q63 | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt, Herbert_Akroyd_Stuart_0222bb.txt | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt | ✅ |
| q64 | United_States_v._Paramount_Pictures__Inc._4aa665.txt, Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | ✅ |
| q65 | Children_s_Mercy_Park_c9283f.txt, Arrowhead_Stadium_eae21e.txt, CommunityAmerica_Ballpark_1e3f7d.txt | Children_s_Mercy_Park_c9283f.txt, CommunityAmerica_Ballpark_1e3f7d.txt | ✅ |
| q66 | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt, The_Informant__376ad6.txt | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt | ✅ |
| q67 | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt, Eucryphia_3801b3.txt | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt | ✅ |
| q68 | Something_There_380707.txt, Paige_O_Hara_4492a0.txt, Something_There_380707.txt | Paige_O_Hara_4492a0.txt, Something_There_380707.txt | ✅ |
| q69 | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt, Laleli_Mosque_c7818f.txt | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt | ✅ |
| q70 | Opry_Mills_53c375.txt, Music_City_Queen_82e4a1.txt, Southaven__Mississippi_7cc376.txt | Music_City_Queen_82e4a1.txt, Opry_Mills_53c375.txt | ✅ |
| q71 | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt, Opry_Mills_53c375.txt | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt | ✅ |
| q72 | James_Fieser_e21429.txt, Berea_College_88c6f9.txt, Azusa_Pacific_University_fcee08.txt | James_Fieser_e21429.txt, Berea_College_88c6f9.txt | ✅ |
| q73 | James_Burke__science_historian__6fe4bf.txt, Connections__TV_series__deb419.txt, Connections__TV_series__deb419.txt | Connections__TV_series__deb419.txt, James_Burke__science_historian__6fe4bf.txt | ✅ |
| q74 | Romeo_87e7c1.txt, Benvolio_848ddf.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt | Romeo_87e7c1.txt, Benvolio_848ddf.txt | ✅ |
| q75 | Addis_Ababa_6c835e.txt, National_Archives_and_Library_of_Ethiopia_6068a1.txt, Ogallala__Nebraska_3578d6.txt | National_Archives_and_Library_of_Ethiopia_6068a1.txt, Addis_Ababa_6c835e.txt | ✅ |
| q76 | Night_Ferry__composition__9f4c6a.txt, Toshi_Ichiyanagi_e50d8d.txt, Symphony_Center_f29c57.txt | Night_Ferry__composition__9f4c6a.txt, Symphony_Center_f29c57.txt | ✅ |
| q77 | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt, Laura_Osnes_4365e8.txt | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt | ✅ |
| q78 | Eucryphia_3801b3.txt, Lepidozamia_4d1d3c.txt, Elatostema_fff908.txt | Lepidozamia_4d1d3c.txt, Eucryphia_3801b3.txt | ✅ |
| q79 | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt, Samoa_086113.txt | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt | ✅ |
| q80 | Kris_Marshall_97e2d7.txt, Death_in_Paradise__TV_series__8a650d.txt, Eddie_Izzard_fb2be0.txt | Death_in_Paradise__TV_series__8a650d.txt, Kris_Marshall_97e2d7.txt | ✅ |
| q81 | EgyptAir_Flight_990_dfd74a.txt, Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt | Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt | ✅ |
| q82 | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt, Oz_the_Great_and_Powerful_510b50.txt | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt | ✅ |
| q83 | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt, Jacques_Sernas_2c77cc.txt | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt | ✅ |
| q84 | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt, Samoa_086113.txt | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt | ✅ |
| q85 | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt, Second_Anglo-Afghan_War_adb0b2.txt | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt | ✅ |
| q86 | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt, Bolton_6ce6c9.txt | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt | ✅ |
| q87 | Hot_air_engine_2ed8e1.txt, Herbert_Akroyd_Stuart_0222bb.txt, George_Cayley_2c8397.txt | George_Cayley_2c8397.txt, Hot_air_engine_2ed8e1.txt | ✅ |
| q88 | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt, Beauty_and_the_Beast__1991_film__d38192.txt | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt | ✅ |
| q89 | Northumbrian_dialect_7334ca.txt, Kingdom_of_Northumbria_4036b8.txt, Kingdom_of_the_Isles_9a036f.txt | Kingdom_of_Northumbria_4036b8.txt, Northumbrian_dialect_7334ca.txt | ✅ |
| q90 | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt, Southaven__Mississippi_7cc376.txt | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt | ✅ |
| q91 | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt | ✅ |
| q92 | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt | ✅ |
| q93 | Oedipus_Rex_d47dfb.txt, Dostoevsky_and_Parricide_f04c2c.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | Dostoevsky_and_Parricide_f04c2c.txt, Oedipus_Rex_d47dfb.txt | ✅ |
| q94 | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt, Chelsea_Peretti_198d48.txt | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt | ✅ |
| q95 | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt, Kunming_08b9a1.txt | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt | ✅ |
| q96 | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt | ✅ |
| q97 | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt, Adnan_Akmal_e863b9.txt | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt | ✅ |
| q98 | Arrowhead_Stadium_eae21e.txt, Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt | Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt | ✅ |
| q99 | Happy_Death_Day_793b37.txt, Jessica_Rothe_59deb4.txt, The_Bye_Bye_Man_906e6d.txt | Jessica_Rothe_59deb4.txt, Happy_Death_Day_793b37.txt | ✅ |
| q100 | Garden_Island_Naval_Chapel_6a3c4e.txt, Royal_Australian_Navy_5e1d16.txt, Samoa_086113.txt | Royal_Australian_Navy_5e1d16.txt, Garden_Island_Naval_Chapel_6a3c4e.txt | ✅ |
| q101 | The_Informant__376ad6.txt, Mark_Whitacre_ccc607.txt, Awake__film__360ee6.txt | Mark_Whitacre_ccc607.txt, The_Informant__376ad6.txt | ✅ |
| q102 | Current_Mood_0317b8.txt, Small_Town_Boy__song__d2fddb.txt, Based_on_a_True_Story..._42790b.txt | Small_Town_Boy__song__d2fddb.txt, Current_Mood_0317b8.txt | ✅ |
| q103 | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt, Impresario_5cc7af.txt | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt | ✅ |
| q104 | Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt, Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt | Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt | ✅ |
| q105 | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt, Tim_Hecker_80429a.txt | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt | ✅ |
| q106 | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt, Adrian_Peterson_1ca55b.txt | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt | ✅ |
| q107 | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt, Goetheanum_5744ff.txt | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt | ✅ |
| q108 | The_Ganymede_Takeover_5feff5.txt, The_Man_in_the_High_Castle_1c064a.txt, Wendell_Berry_4313ab.txt | The_Man_in_the_High_Castle_1c064a.txt, The_Ganymede_Takeover_5feff5.txt | ✅ |
| q109 | Curt_Menefee_0ab3d9.txt, Michael_Strahan_1fa88f.txt, Michael_Strahan_1fa88f.txt | Michael_Strahan_1fa88f.txt, Curt_Menefee_0ab3d9.txt | ✅ |
| q110 | Summer_of_the_Monkeys_0b84ea.txt, William_Allen_White_d1418b.txt, T._R._M._Howard_0bb121.txt | William_Allen_White_d1418b.txt, Summer_of_the_Monkeys_0b84ea.txt | ✅ |
| q111 | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt, Shoba_Chandrasekhar_5285b2.txt | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt | ✅ |
| q112 | Alice_Upside_Down_2fe728.txt, Lucy_Fry_218cb3.txt, Caroline_Carver__actress__7577ed.txt | Lucas_Grabeel_c00cb8.txt, Alice_Upside_Down_2fe728.txt | ✅ |
| q113 | Snowdrop__game_engine__750d41.txt, Tom_Clancy_s_The_Division_3e09f9.txt, Icehouse_pieces_5e75b5.txt | Tom_Clancy_s_The_Division_3e09f9.txt, Snowdrop__game_engine__750d41.txt | ✅ |
| q114 | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt, Tom_Clancy_s_The_Division_3e09f9.txt | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt | ✅ |
| q115 | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt, Jean-Loup_Chrétien_1d1a07.txt | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt | ✅ |
| q116 | Banshee_5e6ebd.txt, VMAQT-1_05fe33.txt, VMAQT-1_05fe33.txt | VMAQT-1_05fe33.txt, Banshee_5e6ebd.txt | ✅ |
| q117 | Barbara_Niven_fbf739.txt, Awake__film__360ee6.txt, Alice_Upside_Down_2fe728.txt | Dead_at_17_273b88.txt, Barbara_Niven_fbf739.txt | ✅ |
| q118 | Bart_the_Fink_77de9b.txt, Krusty_the_Clown_3e3656.txt, Cedric_the_Entertainer_e39d6e.txt | Krusty_the_Clown_3e3656.txt, Bart_the_Fink_77de9b.txt | ✅ |
| q119 | Viaport_Rotterdam_760a19.txt, Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt | Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt | ✅ |
| q120 | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt, Marco_Da_Silva__dancer__777e91.txt | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt | ✅ |
| q121 | Ambrose_Mendy_eddce7.txt, Chris_Eubank_Jr._82fe88.txt, Peter_Schmeichel_0b144a.txt | Chris_Eubank_Jr._82fe88.txt, Ambrose_Mendy_eddce7.txt | ✅ |
| q122 | Allen__amp__Company_Sun_Valley_Conference_083cbf.txt, Rupert_Murdoch_8801f1.txt, Joe_Scarborough_fde209.txt | Rupert_Murdoch_8801f1.txt, Allen__amp__Company_Sun_Valley_Conference_083cbf.txt | ✅ |
| q123 | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt, Raymond_Ochoa_da4d56.txt | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt | ✅ |
| q124 | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt, Phoenix_Television_29103f.txt | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt | ✅ |
| q125 | Patricia_Longo_b7fcef.txt, Graduados_7592c5.txt, Tenerife_c81266.txt | Graduados_7592c5.txt, Patricia_Longo_b7fcef.txt | ✅ |
| q126 | Ogallala_Aquifer_a2b49c.txt, Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt | Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt | ✅ |
| q127 | Blinding_Edge_Pictures_b8de5a.txt, Unbreakable__film__52d8de.txt, Tron_8f60c9.txt | Unbreakable__film__52d8de.txt, Blinding_Edge_Pictures_b8de5a.txt | ✅ |
| q128 | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt, The_Good_Dinosaur_170ac4.txt | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt | ✅ |
| q129 | BraveStarr_8d412d.txt, Celebrity_Home_Entertainment_c01bf6.txt, Gargoyles__TV_series__acb424.txt | Celebrity_Home_Entertainment_c01bf6.txt, BraveStarr_8d412d.txt | ✅ |
| q130 | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt, The_Informant__376ad6.txt | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt | ✅ |
| q131 | Lucy_Fry_218cb3.txt, Jessica_Rothe_59deb4.txt, Nina_Dobrev_05e14f.txt | Mr._Church_ce0d51.txt, Lucy_Fry_218cb3.txt | ✅ |
| q132 | Shoba_Chandrasekhar_5285b2.txt, Ithu_Engal_Neethi_ad89e5.txt, Soha_Ali_Khan_b0280e.txt | Ithu_Engal_Neethi_ad89e5.txt, Shoba_Chandrasekhar_5285b2.txt | ✅ |
| q133 | Official_Ireland_2ed543.txt, Catholic_Church_in_Ireland_77cac6.txt, Catholic_Church_in_Ireland_77cac6.txt | Catholic_Church_in_Ireland_77cac6.txt, Official_Ireland_2ed543.txt | ✅ |
| q134 | Bridge_to_Terabithia__1985_film__1aaa6c.txt, Bedknobs_and_Broomsticks_090f32.txt, Bridge_to_Terabithia__novel__21de92.txt | Bridge_to_Terabithia__novel__21de92.txt, Bridge_to_Terabithia__1985_film__1aaa6c.txt | ✅ |
| q135 | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt, Curt_Menefee_0ab3d9.txt | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt | ✅ |
| q136 | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt, Kasper_Schmeichel_c9da28.txt | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt | ✅ |
| q137 | Atari_Assembler_Editor_a5f9fc.txt, Shepardson_Microsystems_0fa820.txt, Snowdrop__game_engine__750d41.txt | Shepardson_Microsystems_0fa820.txt, Atari_Assembler_Editor_a5f9fc.txt | ✅ |
| q138 | His_Band_and_the_Street_Choir_f5b88a.txt, I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt | I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt | ✅ |
| q139 | Aldosterone_b32476.txt, Aldosterone_b32476.txt, Angiotensin_3f2772.txt | Angiotensin_3f2772.txt, Aldosterone_b32476.txt | ✅ |
| q140 | Nancy_Soderberg_a37451.txt, United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt | United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt | ✅ |
| q141 | Mississippi_University_for_Women_v._Hogan_f450c0.txt, Berghuis_v._Thompkins_c16323.txt, Reynolds_v._Sims_31c83e.txt | Berghuis_v._Thompkins_c16323.txt, Mississippi_University_for_Women_v._Hogan_f450c0.txt | ✅ |
| q142 | AFC_North_8f76d9.txt, AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt | AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt | ✅ |
| q143 | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt | ✅ |
| q144 | Honda_Ballade_4d8914.txt, Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt | Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt | ✅ |
| q145 | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt, The_Monster__song__886456.txt | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt | ✅ |
| q146 | Backstage__magazine__258d77.txt, Celebrity_Home_Entertainment_c01bf6.txt, Christopher_Oscar_Peña_3991f4.txt | Backstage__magazine__258d77.txt, Christopher_Oscar_Peña_3991f4.txt | ✅ |
| q147 | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt, Music_for_Electric_Metronomes_0c93d9.txt | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt | ✅ |
| q148 | Ego_the_Living_Planet_f6f847.txt, Guardians_of_the_Galaxy_Vol._2_b6c488.txt, James_Gunn_9a06f8.txt | Guardians_of_the_Galaxy_Vol._2_b6c488.txt, Ego_the_Living_Planet_f6f847.txt | ✅ |
| q149 | Sponsorship_scandal_5e9ed3.txt, Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt | Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt | ✅ |
| q150 | Conrad_Anker_9d25ca.txt, George_Mallory_ccb70e.txt, Apsley_Cherry-Garrard_ff4e03.txt | George_Mallory_ccb70e.txt, Conrad_Anker_9d25ca.txt | ✅ |
| q151 | Joan_Crawford_5fcb43.txt, The_Duke_Steps_Out_7cc9f2.txt, Erika_Jayne_60e347.txt | The_Duke_Steps_Out_7cc9f2.txt, Joan_Crawford_5fcb43.txt | ✅ |
| q152 | Carol__film__cbfb1f.txt, Guild_of_Music_Supervisors_Awards_7984fd.txt, The_Muppet_Christmas_Carol_81e722.txt | Guild_of_Music_Supervisors_Awards_7984fd.txt, Carol__film__cbfb1f.txt | ✅ |
| q153 | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt | ✅ |
| q154 | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt | ✅ |
| q155 | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt, Goetheanum_5744ff.txt | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt | ✅ |
| q156 | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt, Berea_College_88c6f9.txt | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt | ✅ |
| q157 | Sprawl_trilogy_894d39.txt, Neuromancer_6029bf.txt, The_Man_in_the_High_Castle_1c064a.txt | Neuromancer_6029bf.txt, Sprawl_trilogy_894d39.txt | ✅ |
| q158 | Azad_Hind_Dal_a97cda.txt, Subhas_Chandra_Bose_4b75c6.txt, Second_Anglo-Afghan_War_adb0b2.txt | Subhas_Chandra_Bose_4b75c6.txt, Azad_Hind_Dal_a97cda.txt | ✅ |
| q159 | Eddie_Izzard_fb2be0.txt, Stripped__tour__91c664.txt, Chelsea_Peretti_198d48.txt | Stripped__tour__91c664.txt, Eddie_Izzard_fb2be0.txt | ✅ |
| q160 | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt, Soha_Ali_Khan_b0280e.txt | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt | ✅ |
| q161 | Ecballium_ad2282.txt, Elatostema_fff908.txt, Thalictrum_3222b5.txt | Elatostema_fff908.txt, Ecballium_ad2282.txt | ✅ |
| q162 | Polypodium_71abb6.txt, Aichryson_319a15.txt, Thalictrum_3222b5.txt | Polypodium_71abb6.txt, Aichryson_319a15.txt | ✅ |
| q163 | Adoption_2002_46d793.txt, Adoption_and_Safe_Families_Act_b4c27b.txt, Conscription_in_the_United_States_e920fa.txt | Adoption_and_Safe_Families_Act_b4c27b.txt, Adoption_2002_46d793.txt | ✅ |
| q164 | Crash_Pad_ed95da.txt, Nina_Dobrev_05e14f.txt, The_Lodge__TV_series__b905ce.txt | Nina_Dobrev_05e14f.txt, Crash_Pad_ed95da.txt | ✅ |
| q165 | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt, Their_Lives_eed122.txt | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt | ✅ |
| q166 | Parks_and_Recreation_051519.txt, Ms._Knope_Goes_to_Washington_799ad4.txt, Will__amp__Grace_31b102.txt | Ms._Knope_Goes_to_Washington_799ad4.txt, Parks_and_Recreation_051519.txt | ✅ |
| q167 | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt, Kasper_Schmeichel_c9da28.txt | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt | ✅ |
| q168 | Can_t_Fight_the_Moonlight_caf06b.txt, The_Best_of_LeAnn_Rimes_722f10.txt, Tron_8f60c9.txt | The_Best_of_LeAnn_Rimes_722f10.txt, Can_t_Fight_the_Moonlight_caf06b.txt | ✅ |
| q169 | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt, Lee_Seung-gi_1ed0e1.txt | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt | ✅ |
| q170 | Unbelievable__The_Notorious_B.I.G._song__91da75.txt, Ready_to_Die_c706b2.txt, Ready_to_Die_c706b2.txt | Ready_to_Die_c706b2.txt, Unbelievable__The_Notorious_B.I.G._song__91da75.txt | ✅ |
| q171 | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt, Cedric_the_Entertainer_e39d6e.txt | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt | ✅ |
| q172 | Violin_Sonata_No._5__Beethoven__8f6942.txt, Symphony_No._7__Beethoven__9c3b01.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | Symphony_No._7__Beethoven__9c3b01.txt, Violin_Sonata_No._5__Beethoven__8f6942.txt | ✅ |
| q173 | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt, Mondelez_International_9a57f2.txt | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt | ✅ |
| q174 | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt, Cedric_the_Entertainer_e39d6e.txt | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt | ✅ |
| q175 | Saab_36_108e14.txt, Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt | Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt | ✅ |
| q176 | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt, Pasek_and_Paul_aa5312.txt | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt | ✅ |
| q177 | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt, North_American_T-6_Texan_76936b.txt | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt | ✅ |
| q178 | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt, T._R._M._Howard_0bb121.txt | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt | ✅ |
| q179 | Lucas_Carvalho_940f12.txt, 4___400_metres_relay_8bba84.txt, 4___400_metres_relay_8bba84.txt | 4___400_metres_relay_8bba84.txt, Lucas_Carvalho_940f12.txt | ✅ |
| q180 | Watercliffe_Meadow_Community_Primary_School_1e4bca.txt, Political_correctness_f6d0c6.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt | Political_correctness_f6d0c6.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt | ✅ |
| q181 | Parodia_a9ceb8.txt, Thalictrum_3222b5.txt, Polypodium_71abb6.txt | Thalictrum_3222b5.txt, Parodia_a9ceb8.txt | ✅ |
| q182 | The_Simpsons__season_23__36cc7b.txt, At_Long_Last_Leave_8151f2.txt, Bart_the_Fink_77de9b.txt | At_Long_Last_Leave_8151f2.txt, The_Simpsons__season_23__36cc7b.txt | ✅ |
| q183 | BJ_s_Wholesale_Club_f114e3.txt, US_Vision_f98d42.txt, US_Vision_f98d42.txt | US_Vision_f98d42.txt, BJ_s_Wholesale_Club_f114e3.txt | ✅ |
| q184 | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt, Lochmere_Archeological_District_2ae282.txt | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt | ✅ |
| q185 | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt, Southaven__Mississippi_7cc376.txt | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt | ✅ |
| q186 | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt, Samantha_Cristoforetti_fc1bcd.txt | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt | ✅ |
| q187 | Flynn_Rider_37adf2.txt, Zachary_Levi_9c18c7.txt, Lucas_Grabeel_c00cb8.txt | Zachary_Levi_9c18c7.txt, Flynn_Rider_37adf2.txt | ✅ |
| q188 | Vacation_with_Derek_da8ab6.txt, Life_with_Derek_547274.txt, Cressida_Bonas_5c3c57.txt | Life_with_Derek_547274.txt, Vacation_with_Derek_da8ab6.txt | ✅ |
| q189 | Jin_Jing_fa2e08.txt, Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt | Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt | ✅ |
| q190 | The_Worst_Journey_in_the_World_1fd01f.txt, Apsley_Cherry-Garrard_ff4e03.txt, Apsley_Cherry-Garrard_ff4e03.txt | Apsley_Cherry-Garrard_ff4e03.txt, The_Worst_Journey_in_the_World_1fd01f.txt | ✅ |
| q191 | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt, James_Fieser_e21429.txt | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | ✅ |
| q192 | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, The_Prince_and_Me_253bec.txt, Lee_Seung-gi_1ed0e1.txt | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, Lee_Seung-gi_1ed0e1.txt | ✅ |
| q193 | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt | ✅ |
| q194 | Thomas_Doherty__actor__7caf88.txt, The_Lodge__TV_series__b905ce.txt, Life_with_Derek_547274.txt | The_Lodge__TV_series__b905ce.txt, Thomas_Doherty__actor__7caf88.txt | ✅ |
| q195 | Chelsea_Peretti_198d48.txt, Brooklyn_Nine-Nine_eec7c9.txt, Larry_Drake_85028c.txt | Brooklyn_Nine-Nine_eec7c9.txt, Chelsea_Peretti_198d48.txt | ✅ |
| q196 | Mark_Gaudet_0f4aa3.txt, Jan_Axel_Blomberg_a42602.txt, Mark_Lindsay_c8bd25.txt | Jan_Axel_Blomberg_a42602.txt, Mark_Gaudet_0f4aa3.txt | ✅ |
| q197 | Miami_Canal_161997.txt, Dundee_Canal_05ea0b.txt, Richard_Hornsby__amp__Sons_fb744c.txt | Dundee_Canal_05ea0b.txt, Miami_Canal_161997.txt | ✅ |
| q198 | Tron_8f60c9.txt, The_Million_Dollar_Duck_d1d45c.txt, Gryphon__film__f811a3.txt | The_Million_Dollar_Duck_d1d45c.txt, Tron_8f60c9.txt | ✅ |
| q199 | Beauty_and_the_Beast__franchise__7a780a.txt, Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt | Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt | ✅ |
| q200 | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt, Peter_Schmeichel_0b144a.txt | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt | ✅ |

### Config: top5 (top_k=5)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt, Eski_Imaret_Mosque_b4d62c.txt (+2) | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt | ✅ |
| q2 | Random_House_Tower_2b6a2a.txt, 888_7th_Avenue_f091f0.txt, 888_7th_Avenue_f091f0.txt (+2) | 888_7th_Avenue_f091f0.txt, Random_House_Tower_2b6a2a.txt | ✅ |
| q3 | Alex_Ferguson_6962e9.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt, Peter_Schmeichel_0b144a.txt (+2) | 1995_96_Manchester_United_F.C._season_3f9756.txt, Alex_Ferguson_6962e9.txt | ✅ |
| q4 | Apple_Remote_fcaffa.txt, Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt (+2) | Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt | ✅ |
| q5 | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt, Charles_Nungesser_67ba12.txt (+2) | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt | ✅ |
| q6 | Henry_J._Kaiser_53c448.txt, Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt (+2) | Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt | ✅ |
| q7 | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt, Jean-Loup_Chrétien_1d1a07.txt (+2) | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt | ✅ |
| q8 | Freakonomics__film__249214.txt, In_the_Realm_of_the_Hackers_367f43.txt, Connections__TV_series__deb419.txt (+2) | In_the_Realm_of_the_Hackers_367f43.txt, Freakonomics__film__249214.txt | ✅ |
| q9 | Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt (+2) | Socialist_Revolutionary_Party_68737a.txt, Russian_Civil_War_5a6752.txt | ✅ |
| q10 | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt, Ogallala__Nebraska_3578d6.txt (+2) | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt | ✅ |
| q11 | Giuseppe_Arimondi_0fc7eb.txt, Battle_of_Adwa_1fa890.txt, Addis_Ababa_6c835e.txt (+2) | Battle_of_Adwa_1fa890.txt, Giuseppe_Arimondi_0fc7eb.txt | ✅ |
| q12 | Dirleton_Castle_a042b8.txt, Yellowcraigs_0c745c.txt, Kingdom_of_Northumbria_4036b8.txt (+2) | Yellowcraigs_0c745c.txt, Dirleton_Castle_a042b8.txt | ✅ |
| q13 | English_Electric_Canberra_becfbe.txt, Avro_Vulcan_7aa981.txt, English_Electric_Canberra_becfbe.txt (+2) | English_Electric_Canberra_becfbe.txt, No._2_Squadron_RAAF_a75482.txt | ✅ |
| q14 | Euromarché_085339.txt, Carrefour_a2a75b.txt, Maxeda_f76dfe.txt (+2) | Euromarché_085339.txt, Carrefour_a2a75b.txt | ✅ |
| q15 | Delirium__Ellie_Goulding_album__5bb0cb.txt, On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Erika_Jayne_60e347.txt (+2) | On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Delirium__Ellie_Goulding_album__5bb0cb.txt | ✅ |
| q16 | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt, The_Legend_of_Korra_86cdc2.txt (+2) | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt | ✅ |
| q17 | Oranjegekte_8248a5.txt, Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt (+2) | Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt | ✅ |
| q18 | Tromeo_and_Juliet_86fab3.txt, James_Gunn_9a06f8.txt, Romeo_87e7c1.txt (+2) | James_Gunn_9a06f8.txt, Tromeo_and_Juliet_86fab3.txt | ✅ |
| q19 | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt, Bob_Seger_69d05d.txt (+2) | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt | ✅ |
| q20 | Rostker_v._Goldberg_b61238.txt, Conscription_in_the_United_States_e920fa.txt, Conscription_in_the_United_States_e920fa.txt (+2) | Conscription_in_the_United_States_e920fa.txt, Rostker_v._Goldberg_b61238.txt | ✅ |
| q21 | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt, Orange_Julius_e106c8.txt (+2) | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt | ✅ |
| q22 | Their_Lives_eed122.txt, Monica_Lewinsky_7bb6c2.txt, Nancy_Soderberg_a37451.txt (+2) | Monica_Lewinsky_7bb6c2.txt, Their_Lives_eed122.txt | ✅ |
| q23 | Teide_National_Park_aaf674.txt, Garajonay_National_Park_97f362.txt, Hatton_Castle__Angus_eb96ea.txt (+2) | Garajonay_National_Park_97f362.txt, Teide_National_Park_aaf674.txt | ✅ |
| q24 | Andrew_Jaspan_f6dc15.txt, Andrew_Jaspan_f6dc15.txt, The_Conversation__website__724191.txt (+2) | The_Conversation__website__724191.txt, Andrew_Jaspan_f6dc15.txt | ✅ |
| q25 | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt, The_Conversation__website__724191.txt (+2) | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt | ✅ |
| q26 | Tysons_Galleria_ead975.txt, Oldham_County__Kentucky_34ed53.txt, Ogallala__Nebraska_3578d6.txt (+2) | Tysons_Galleria_ead975.txt, McLean__Virginia_45ec68.txt | ✅ |
| q27 | My_Eyes__Blake_Shelton_song__fa840b.txt, Based_on_a_True_Story..._42790b.txt, The_Best_of_LeAnn_Rimes_722f10.txt (+2) | Based_on_a_True_Story..._42790b.txt, My_Eyes__Blake_Shelton_song__fa840b.txt | ✅ |
| q28 | Caroline_Carver__actress__7577ed.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Carol__film__cbfb1f.txt (+2) | The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Caroline_Carver__actress__7577ed.txt | ✅ |
| q29 | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt, Wilton_Mall_66971e.txt (+2) | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt | ✅ |
| q30 | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt, Jessica_Rothe_59deb4.txt (+2) | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt | ✅ |
| q31 | Mummulgum_1292e7.txt, Casino__New_South_Wales_8c85c5.txt, Ogallala__Nebraska_3578d6.txt (+2) | Casino__New_South_Wales_8c85c5.txt, Mummulgum_1292e7.txt | ✅ |
| q32 | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt, Sacred_Planet_26fd7b.txt (+2) | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt | ✅ |
| q33 | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt, James_Gunn_9a06f8.txt (+2) | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt | ✅ |
| q34 | Roberta_Vinci_f714c8.txt, Jorge_Lozano_1baa7a.txt, Jorge_Lozano_1baa7a.txt (+2) | Jorge_Lozano_1baa7a.txt, Roberta_Vinci_f714c8.txt | ✅ |
| q35 | Marco_Da_Silva__dancer__777e91.txt, Erika_Jayne_60e347.txt, Cressida_Bonas_5c3c57.txt (+2) | Erika_Jayne_60e347.txt, Marco_Da_Silva__dancer__777e91.txt | ✅ |
| q36 | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt (+2) | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt | ✅ |
| q37 | Kingdom_of_the_Isles_9a036f.txt, Kingdom_of_the_Isles_9a036f.txt, Aonghus_Mór_b5e643.txt (+2) | Aonghus_Mór_b5e643.txt, Kingdom_of_the_Isles_9a036f.txt | ✅ |
| q38 | Bruce_Spizer_c78d7e.txt, Bob_Seger_69d05d.txt, The_Beatles_c9e770.txt (+2) | The_Beatles_c9e770.txt, Bruce_Spizer_c78d7e.txt | ✅ |
| q39 | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt, Wayne_Garland_02f688.txt (+2) | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt | ✅ |
| q40 | Argand_lamp_563fb5.txt, Lewis_lamp_ddcc57.txt, Lewis_lamp_ddcc57.txt (+2) | Lewis_lamp_ddcc57.txt, Argand_lamp_563fb5.txt | ✅ |
| q41 | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt, Kathy_Sullivan__Australian_politician__a2272b.txt (+2) | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt | ✅ |
| q42 | Bishop_Carroll_Catholic_High_School_5aa312.txt, Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+2) | Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt | ✅ |
| q43 | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt, United_States_presidential_election__2016_363dfd.txt (+2) | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt | ✅ |
| q44 | Southaven__Mississippi_7cc376.txt, Memphis_Hustle_71bf0b.txt, Ogallala__Nebraska_3578d6.txt (+2) | Memphis_Hustle_71bf0b.txt, Southaven__Mississippi_7cc376.txt | ✅ |
| q45 | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt, Something_There_380707.txt (+2) | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt | ✅ |
| q46 | Albertina_eec665.txt, Hanna_Varis_e45643.txt, Hanna_Varis_e45643.txt (+2) | Albertina_eec665.txt, Hanna_Varis_e45643.txt | ✅ |
| q47 | Hatton_Hill_027377.txt, Hatton_Castle__Angus_eb96ea.txt, Hatton_Castle__Angus_eb96ea.txt (+2) | Hatton_Castle__Angus_eb96ea.txt, Hatton_Hill_027377.txt | ✅ |
| q48 | The_Legend_of_Korra_86cdc2.txt, Kuvira_4fdaa0.txt, Gargoyles__TV_series__acb424.txt (+2) | Kuvira_4fdaa0.txt, The_Legend_of_Korra_86cdc2.txt | ✅ |
| q49 | The_Five_Obstructions_18f018.txt, The_Importance_of_Being_Icelandic_086a7b.txt, The_Importance_of_Being_Icelandic_086a7b.txt (+2) | The_Importance_of_Being_Icelandic_086a7b.txt, The_Five_Obstructions_18f018.txt | ✅ |
| q50 | Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt, Will__amp__Grace_31b102.txt, The_Legend_of_Korra_86cdc2.txt (+2) | Will__amp__Grace_31b102.txt, Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt | ✅ |
| q51 | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt, Kohlberg_Kravis_Roberts_728df7.txt (+2) | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt | ✅ |
| q52 | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt, Wendell_Berry_4313ab.txt (+2) | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt | ✅ |
| q53 | 1920__film__db5e98.txt, Soha_Ali_Khan_b0280e.txt, 1920__film_series__78d752.txt (+2) | 1920__film_series__78d752.txt, 1920__film__db5e98.txt | ✅ |
| q54 | 71st_Golden_Globe_Awards_8dbed6.txt, 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt (+2) | 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt | ✅ |
| q55 | Charles_Hastings_Judd_409cf1.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt, Kalākaua_055843.txt (+2) | Charles_Hastings_Judd_409cf1.txt, Kalākaua_055843.txt | ✅ |
| q56 | Armie_Hammer_7c1778.txt, The_Polar_Bears_b06086.txt, Eddie_Izzard_fb2be0.txt (+2) | The_Polar_Bears_b06086.txt, Armie_Hammer_7c1778.txt | ✅ |
| q57 | 712_Fifth_Avenue_4807f9.txt, Manhattan_Life_Insurance_Building_fb6468.txt, Random_House_Tower_2b6a2a.txt (+2) | Manhattan_Life_Insurance_Building_fb6468.txt, 712_Fifth_Avenue_4807f9.txt | ✅ |
| q58 | Tenerife_c81266.txt, Gerald_Reive_c5cd23.txt, Samoa_086113.txt (+2) | Samoa_086113.txt, Gerald_Reive_c5cd23.txt | ✅ |
| q59 | Tecumseh_bcf68b.txt, Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt (+2) | Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt | ✅ |
| q60 | Samuel_Sim_07d71f.txt, Tromeo_and_Juliet_86fab3.txt, Bedknobs_and_Broomsticks_090f32.txt (+2) | Samuel_Sim_07d71f.txt, Awake__film__360ee6.txt | ✅ |
| q61 | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt, Øresund_Region_549ba5.txt (+2) | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt | ✅ |
| q62 | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt, Larry_Drake_85028c.txt (+2) | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt | ✅ |
| q63 | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt, Herbert_Akroyd_Stuart_0222bb.txt (+2) | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt | ✅ |
| q64 | United_States_v._Paramount_Pictures__Inc._4aa665.txt, Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+2) | Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | ✅ |
| q65 | Children_s_Mercy_Park_c9283f.txt, Arrowhead_Stadium_eae21e.txt, CommunityAmerica_Ballpark_1e3f7d.txt (+2) | Children_s_Mercy_Park_c9283f.txt, CommunityAmerica_Ballpark_1e3f7d.txt | ✅ |
| q66 | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt, The_Informant__376ad6.txt (+2) | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt | ✅ |
| q67 | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt, Eucryphia_3801b3.txt (+2) | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt | ✅ |
| q68 | Something_There_380707.txt, Paige_O_Hara_4492a0.txt, Something_There_380707.txt (+2) | Paige_O_Hara_4492a0.txt, Something_There_380707.txt | ✅ |
| q69 | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt, Laleli_Mosque_c7818f.txt (+2) | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt | ✅ |
| q70 | Opry_Mills_53c375.txt, Music_City_Queen_82e4a1.txt, Southaven__Mississippi_7cc376.txt (+2) | Music_City_Queen_82e4a1.txt, Opry_Mills_53c375.txt | ✅ |
| q71 | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt, Opry_Mills_53c375.txt (+2) | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt | ✅ |
| q72 | James_Fieser_e21429.txt, Berea_College_88c6f9.txt, Azusa_Pacific_University_fcee08.txt (+2) | James_Fieser_e21429.txt, Berea_College_88c6f9.txt | ✅ |
| q73 | James_Burke__science_historian__6fe4bf.txt, Connections__TV_series__deb419.txt, Connections__TV_series__deb419.txt (+2) | Connections__TV_series__deb419.txt, James_Burke__science_historian__6fe4bf.txt | ✅ |
| q74 | Romeo_87e7c1.txt, Benvolio_848ddf.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt (+2) | Romeo_87e7c1.txt, Benvolio_848ddf.txt | ✅ |
| q75 | Addis_Ababa_6c835e.txt, National_Archives_and_Library_of_Ethiopia_6068a1.txt, Ogallala__Nebraska_3578d6.txt (+2) | National_Archives_and_Library_of_Ethiopia_6068a1.txt, Addis_Ababa_6c835e.txt | ✅ |
| q76 | Night_Ferry__composition__9f4c6a.txt, Toshi_Ichiyanagi_e50d8d.txt, Symphony_Center_f29c57.txt (+2) | Night_Ferry__composition__9f4c6a.txt, Symphony_Center_f29c57.txt | ✅ |
| q77 | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt, Laura_Osnes_4365e8.txt (+2) | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt | ✅ |
| q78 | Eucryphia_3801b3.txt, Lepidozamia_4d1d3c.txt, Elatostema_fff908.txt (+2) | Lepidozamia_4d1d3c.txt, Eucryphia_3801b3.txt | ✅ |
| q79 | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt, Samoa_086113.txt (+2) | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt | ✅ |
| q80 | Kris_Marshall_97e2d7.txt, Death_in_Paradise__TV_series__8a650d.txt, Eddie_Izzard_fb2be0.txt (+2) | Death_in_Paradise__TV_series__8a650d.txt, Kris_Marshall_97e2d7.txt | ✅ |
| q81 | EgyptAir_Flight_990_dfd74a.txt, Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt (+2) | Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt | ✅ |
| q82 | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt, Oz_the_Great_and_Powerful_510b50.txt (+2) | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt | ✅ |
| q83 | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt, Jacques_Sernas_2c77cc.txt (+2) | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt | ✅ |
| q84 | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt, Samoa_086113.txt (+2) | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt | ✅ |
| q85 | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt, Second_Anglo-Afghan_War_adb0b2.txt (+2) | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt | ✅ |
| q86 | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt, Bolton_6ce6c9.txt (+2) | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt | ✅ |
| q87 | Hot_air_engine_2ed8e1.txt, Herbert_Akroyd_Stuart_0222bb.txt, George_Cayley_2c8397.txt (+2) | George_Cayley_2c8397.txt, Hot_air_engine_2ed8e1.txt | ✅ |
| q88 | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt, Beauty_and_the_Beast__1991_film__d38192.txt (+2) | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt | ✅ |
| q89 | Northumbrian_dialect_7334ca.txt, Kingdom_of_Northumbria_4036b8.txt, Kingdom_of_the_Isles_9a036f.txt (+2) | Kingdom_of_Northumbria_4036b8.txt, Northumbrian_dialect_7334ca.txt | ✅ |
| q90 | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt, Southaven__Mississippi_7cc376.txt (+2) | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt | ✅ |
| q91 | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt (+2) | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt | ✅ |
| q92 | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+2) | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt | ✅ |
| q93 | Oedipus_Rex_d47dfb.txt, Dostoevsky_and_Parricide_f04c2c.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+2) | Dostoevsky_and_Parricide_f04c2c.txt, Oedipus_Rex_d47dfb.txt | ✅ |
| q94 | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt, Chelsea_Peretti_198d48.txt (+2) | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt | ✅ |
| q95 | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt, Kunming_08b9a1.txt (+2) | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt | ✅ |
| q96 | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+2) | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt | ✅ |
| q97 | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt, Adnan_Akmal_e863b9.txt (+2) | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt | ✅ |
| q98 | Arrowhead_Stadium_eae21e.txt, Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt (+2) | Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt | ✅ |
| q99 | Happy_Death_Day_793b37.txt, Jessica_Rothe_59deb4.txt, The_Bye_Bye_Man_906e6d.txt (+2) | Jessica_Rothe_59deb4.txt, Happy_Death_Day_793b37.txt | ✅ |
| q100 | Garden_Island_Naval_Chapel_6a3c4e.txt, Royal_Australian_Navy_5e1d16.txt, Samoa_086113.txt (+2) | Royal_Australian_Navy_5e1d16.txt, Garden_Island_Naval_Chapel_6a3c4e.txt | ✅ |
| q101 | The_Informant__376ad6.txt, Mark_Whitacre_ccc607.txt, Awake__film__360ee6.txt (+2) | Mark_Whitacre_ccc607.txt, The_Informant__376ad6.txt | ✅ |
| q102 | Current_Mood_0317b8.txt, Small_Town_Boy__song__d2fddb.txt, Based_on_a_True_Story..._42790b.txt (+2) | Small_Town_Boy__song__d2fddb.txt, Current_Mood_0317b8.txt | ✅ |
| q103 | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt, Impresario_5cc7af.txt (+2) | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt | ✅ |
| q104 | Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt, Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt (+2) | Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt | ✅ |
| q105 | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt, Tim_Hecker_80429a.txt (+2) | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt | ✅ |
| q106 | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt, Adrian_Peterson_1ca55b.txt (+2) | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt | ✅ |
| q107 | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt, Goetheanum_5744ff.txt (+2) | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt | ✅ |
| q108 | The_Ganymede_Takeover_5feff5.txt, The_Man_in_the_High_Castle_1c064a.txt, Wendell_Berry_4313ab.txt (+2) | The_Man_in_the_High_Castle_1c064a.txt, The_Ganymede_Takeover_5feff5.txt | ✅ |
| q109 | Curt_Menefee_0ab3d9.txt, Michael_Strahan_1fa88f.txt, Michael_Strahan_1fa88f.txt (+2) | Michael_Strahan_1fa88f.txt, Curt_Menefee_0ab3d9.txt | ✅ |
| q110 | Summer_of_the_Monkeys_0b84ea.txt, William_Allen_White_d1418b.txt, T._R._M._Howard_0bb121.txt (+2) | William_Allen_White_d1418b.txt, Summer_of_the_Monkeys_0b84ea.txt | ✅ |
| q111 | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt, Shoba_Chandrasekhar_5285b2.txt (+2) | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt | ✅ |
| q112 | Alice_Upside_Down_2fe728.txt, Lucy_Fry_218cb3.txt, Caroline_Carver__actress__7577ed.txt (+2) | Lucas_Grabeel_c00cb8.txt, Alice_Upside_Down_2fe728.txt | ✅ |
| q113 | Snowdrop__game_engine__750d41.txt, Tom_Clancy_s_The_Division_3e09f9.txt, Icehouse_pieces_5e75b5.txt (+2) | Tom_Clancy_s_The_Division_3e09f9.txt, Snowdrop__game_engine__750d41.txt | ✅ |
| q114 | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt, Tom_Clancy_s_The_Division_3e09f9.txt (+2) | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt | ✅ |
| q115 | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt, Jean-Loup_Chrétien_1d1a07.txt (+2) | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt | ✅ |
| q116 | Banshee_5e6ebd.txt, VMAQT-1_05fe33.txt, VMAQT-1_05fe33.txt (+2) | VMAQT-1_05fe33.txt, Banshee_5e6ebd.txt | ✅ |
| q117 | Barbara_Niven_fbf739.txt, Awake__film__360ee6.txt, Alice_Upside_Down_2fe728.txt (+2) | Dead_at_17_273b88.txt, Barbara_Niven_fbf739.txt | ✅ |
| q118 | Bart_the_Fink_77de9b.txt, Krusty_the_Clown_3e3656.txt, Cedric_the_Entertainer_e39d6e.txt (+2) | Krusty_the_Clown_3e3656.txt, Bart_the_Fink_77de9b.txt | ✅ |
| q119 | Viaport_Rotterdam_760a19.txt, Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt (+2) | Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt | ✅ |
| q120 | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt, Marco_Da_Silva__dancer__777e91.txt (+2) | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt | ✅ |
| q121 | Ambrose_Mendy_eddce7.txt, Chris_Eubank_Jr._82fe88.txt, Peter_Schmeichel_0b144a.txt (+2) | Chris_Eubank_Jr._82fe88.txt, Ambrose_Mendy_eddce7.txt | ✅ |
| q122 | Allen__amp__Company_Sun_Valley_Conference_083cbf.txt, Rupert_Murdoch_8801f1.txt, Joe_Scarborough_fde209.txt (+2) | Rupert_Murdoch_8801f1.txt, Allen__amp__Company_Sun_Valley_Conference_083cbf.txt | ✅ |
| q123 | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt, Raymond_Ochoa_da4d56.txt (+2) | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt | ✅ |
| q124 | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt, Phoenix_Television_29103f.txt (+2) | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt | ✅ |
| q125 | Patricia_Longo_b7fcef.txt, Graduados_7592c5.txt, Tenerife_c81266.txt (+2) | Graduados_7592c5.txt, Patricia_Longo_b7fcef.txt | ✅ |
| q126 | Ogallala_Aquifer_a2b49c.txt, Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt (+2) | Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt | ✅ |
| q127 | Blinding_Edge_Pictures_b8de5a.txt, Unbreakable__film__52d8de.txt, Tron_8f60c9.txt (+2) | Unbreakable__film__52d8de.txt, Blinding_Edge_Pictures_b8de5a.txt | ✅ |
| q128 | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt, The_Good_Dinosaur_170ac4.txt (+2) | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt | ✅ |
| q129 | BraveStarr_8d412d.txt, Celebrity_Home_Entertainment_c01bf6.txt, Gargoyles__TV_series__acb424.txt (+2) | Celebrity_Home_Entertainment_c01bf6.txt, BraveStarr_8d412d.txt | ✅ |
| q130 | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt, The_Informant__376ad6.txt (+2) | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt | ✅ |
| q131 | Lucy_Fry_218cb3.txt, Jessica_Rothe_59deb4.txt, Nina_Dobrev_05e14f.txt (+2) | Mr._Church_ce0d51.txt, Lucy_Fry_218cb3.txt | ✅ |
| q132 | Shoba_Chandrasekhar_5285b2.txt, Ithu_Engal_Neethi_ad89e5.txt, Soha_Ali_Khan_b0280e.txt (+2) | Ithu_Engal_Neethi_ad89e5.txt, Shoba_Chandrasekhar_5285b2.txt | ✅ |
| q133 | Official_Ireland_2ed543.txt, Catholic_Church_in_Ireland_77cac6.txt, Catholic_Church_in_Ireland_77cac6.txt (+2) | Catholic_Church_in_Ireland_77cac6.txt, Official_Ireland_2ed543.txt | ✅ |
| q134 | Bridge_to_Terabithia__1985_film__1aaa6c.txt, Bedknobs_and_Broomsticks_090f32.txt, Bridge_to_Terabithia__novel__21de92.txt (+2) | Bridge_to_Terabithia__novel__21de92.txt, Bridge_to_Terabithia__1985_film__1aaa6c.txt | ✅ |
| q135 | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt, Curt_Menefee_0ab3d9.txt (+2) | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt | ✅ |
| q136 | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt, Kasper_Schmeichel_c9da28.txt (+2) | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt | ✅ |
| q137 | Atari_Assembler_Editor_a5f9fc.txt, Shepardson_Microsystems_0fa820.txt, Snowdrop__game_engine__750d41.txt (+2) | Shepardson_Microsystems_0fa820.txt, Atari_Assembler_Editor_a5f9fc.txt | ✅ |
| q138 | His_Band_and_the_Street_Choir_f5b88a.txt, I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt (+2) | I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt | ✅ |
| q139 | Aldosterone_b32476.txt, Aldosterone_b32476.txt, Angiotensin_3f2772.txt (+2) | Angiotensin_3f2772.txt, Aldosterone_b32476.txt | ✅ |
| q140 | Nancy_Soderberg_a37451.txt, United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt (+2) | United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt | ✅ |
| q141 | Mississippi_University_for_Women_v._Hogan_f450c0.txt, Berghuis_v._Thompkins_c16323.txt, Reynolds_v._Sims_31c83e.txt (+2) | Berghuis_v._Thompkins_c16323.txt, Mississippi_University_for_Women_v._Hogan_f450c0.txt | ✅ |
| q142 | AFC_North_8f76d9.txt, AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt (+2) | AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt | ✅ |
| q143 | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt (+2) | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt | ✅ |
| q144 | Honda_Ballade_4d8914.txt, Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt (+2) | Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt | ✅ |
| q145 | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt, The_Monster__song__886456.txt (+2) | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt | ✅ |
| q146 | Backstage__magazine__258d77.txt, Celebrity_Home_Entertainment_c01bf6.txt, Christopher_Oscar_Peña_3991f4.txt (+2) | Backstage__magazine__258d77.txt, Christopher_Oscar_Peña_3991f4.txt | ✅ |
| q147 | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt, Music_for_Electric_Metronomes_0c93d9.txt (+2) | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt | ✅ |
| q148 | Ego_the_Living_Planet_f6f847.txt, Guardians_of_the_Galaxy_Vol._2_b6c488.txt, James_Gunn_9a06f8.txt (+2) | Guardians_of_the_Galaxy_Vol._2_b6c488.txt, Ego_the_Living_Planet_f6f847.txt | ✅ |
| q149 | Sponsorship_scandal_5e9ed3.txt, Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt (+2) | Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt | ✅ |
| q150 | Conrad_Anker_9d25ca.txt, George_Mallory_ccb70e.txt, Apsley_Cherry-Garrard_ff4e03.txt (+2) | George_Mallory_ccb70e.txt, Conrad_Anker_9d25ca.txt | ✅ |
| q151 | Joan_Crawford_5fcb43.txt, The_Duke_Steps_Out_7cc9f2.txt, Erika_Jayne_60e347.txt (+2) | The_Duke_Steps_Out_7cc9f2.txt, Joan_Crawford_5fcb43.txt | ✅ |
| q152 | Carol__film__cbfb1f.txt, Guild_of_Music_Supervisors_Awards_7984fd.txt, The_Muppet_Christmas_Carol_81e722.txt (+2) | Guild_of_Music_Supervisors_Awards_7984fd.txt, Carol__film__cbfb1f.txt | ✅ |
| q153 | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+2) | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt | ✅ |
| q154 | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt (+2) | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt | ✅ |
| q155 | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt, Goetheanum_5744ff.txt (+2) | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt | ✅ |
| q156 | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt, Berea_College_88c6f9.txt (+2) | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt | ✅ |
| q157 | Sprawl_trilogy_894d39.txt, Neuromancer_6029bf.txt, The_Man_in_the_High_Castle_1c064a.txt (+2) | Neuromancer_6029bf.txt, Sprawl_trilogy_894d39.txt | ✅ |
| q158 | Azad_Hind_Dal_a97cda.txt, Subhas_Chandra_Bose_4b75c6.txt, Second_Anglo-Afghan_War_adb0b2.txt (+2) | Subhas_Chandra_Bose_4b75c6.txt, Azad_Hind_Dal_a97cda.txt | ✅ |
| q159 | Eddie_Izzard_fb2be0.txt, Stripped__tour__91c664.txt, Chelsea_Peretti_198d48.txt (+2) | Stripped__tour__91c664.txt, Eddie_Izzard_fb2be0.txt | ✅ |
| q160 | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt, Soha_Ali_Khan_b0280e.txt (+2) | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt | ✅ |
| q161 | Ecballium_ad2282.txt, Elatostema_fff908.txt, Thalictrum_3222b5.txt (+2) | Elatostema_fff908.txt, Ecballium_ad2282.txt | ✅ |
| q162 | Polypodium_71abb6.txt, Aichryson_319a15.txt, Thalictrum_3222b5.txt (+2) | Polypodium_71abb6.txt, Aichryson_319a15.txt | ✅ |
| q163 | Adoption_2002_46d793.txt, Adoption_and_Safe_Families_Act_b4c27b.txt, Conscription_in_the_United_States_e920fa.txt (+2) | Adoption_and_Safe_Families_Act_b4c27b.txt, Adoption_2002_46d793.txt | ✅ |
| q164 | Crash_Pad_ed95da.txt, Nina_Dobrev_05e14f.txt, The_Lodge__TV_series__b905ce.txt (+2) | Nina_Dobrev_05e14f.txt, Crash_Pad_ed95da.txt | ✅ |
| q165 | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt, Their_Lives_eed122.txt (+2) | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt | ✅ |
| q166 | Parks_and_Recreation_051519.txt, Ms._Knope_Goes_to_Washington_799ad4.txt, Will__amp__Grace_31b102.txt (+2) | Ms._Knope_Goes_to_Washington_799ad4.txt, Parks_and_Recreation_051519.txt | ✅ |
| q167 | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt, Kasper_Schmeichel_c9da28.txt (+2) | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt | ✅ |
| q168 | Can_t_Fight_the_Moonlight_caf06b.txt, The_Best_of_LeAnn_Rimes_722f10.txt, Tron_8f60c9.txt (+2) | The_Best_of_LeAnn_Rimes_722f10.txt, Can_t_Fight_the_Moonlight_caf06b.txt | ✅ |
| q169 | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt, Lee_Seung-gi_1ed0e1.txt (+2) | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt | ✅ |
| q170 | Unbelievable__The_Notorious_B.I.G._song__91da75.txt, Ready_to_Die_c706b2.txt, Ready_to_Die_c706b2.txt (+2) | Ready_to_Die_c706b2.txt, Unbelievable__The_Notorious_B.I.G._song__91da75.txt | ✅ |
| q171 | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt, Cedric_the_Entertainer_e39d6e.txt (+2) | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt | ✅ |
| q172 | Violin_Sonata_No._5__Beethoven__8f6942.txt, Symphony_No._7__Beethoven__9c3b01.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+2) | Symphony_No._7__Beethoven__9c3b01.txt, Violin_Sonata_No._5__Beethoven__8f6942.txt | ✅ |
| q173 | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt, Mondelez_International_9a57f2.txt (+2) | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt | ✅ |
| q174 | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt, Cedric_the_Entertainer_e39d6e.txt (+2) | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt | ✅ |
| q175 | Saab_36_108e14.txt, Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt (+2) | Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt | ✅ |
| q176 | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt, Pasek_and_Paul_aa5312.txt (+2) | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt | ✅ |
| q177 | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt, North_American_T-6_Texan_76936b.txt (+2) | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt | ✅ |
| q178 | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt, T._R._M._Howard_0bb121.txt (+2) | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt | ✅ |
| q179 | Lucas_Carvalho_940f12.txt, 4___400_metres_relay_8bba84.txt, 4___400_metres_relay_8bba84.txt (+2) | 4___400_metres_relay_8bba84.txt, Lucas_Carvalho_940f12.txt | ✅ |
| q180 | Watercliffe_Meadow_Community_Primary_School_1e4bca.txt, Political_correctness_f6d0c6.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+2) | Political_correctness_f6d0c6.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt | ✅ |
| q181 | Parodia_a9ceb8.txt, Thalictrum_3222b5.txt, Polypodium_71abb6.txt (+2) | Thalictrum_3222b5.txt, Parodia_a9ceb8.txt | ✅ |
| q182 | The_Simpsons__season_23__36cc7b.txt, At_Long_Last_Leave_8151f2.txt, Bart_the_Fink_77de9b.txt (+2) | At_Long_Last_Leave_8151f2.txt, The_Simpsons__season_23__36cc7b.txt | ✅ |
| q183 | BJ_s_Wholesale_Club_f114e3.txt, US_Vision_f98d42.txt, US_Vision_f98d42.txt (+2) | US_Vision_f98d42.txt, BJ_s_Wholesale_Club_f114e3.txt | ✅ |
| q184 | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt, Lochmere_Archeological_District_2ae282.txt (+2) | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt | ✅ |
| q185 | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt, Southaven__Mississippi_7cc376.txt (+2) | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt | ✅ |
| q186 | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt, Samantha_Cristoforetti_fc1bcd.txt (+2) | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt | ✅ |
| q187 | Flynn_Rider_37adf2.txt, Zachary_Levi_9c18c7.txt, Lucas_Grabeel_c00cb8.txt (+2) | Zachary_Levi_9c18c7.txt, Flynn_Rider_37adf2.txt | ✅ |
| q188 | Vacation_with_Derek_da8ab6.txt, Life_with_Derek_547274.txt, Cressida_Bonas_5c3c57.txt (+2) | Life_with_Derek_547274.txt, Vacation_with_Derek_da8ab6.txt | ✅ |
| q189 | Jin_Jing_fa2e08.txt, Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt (+2) | Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt | ✅ |
| q190 | The_Worst_Journey_in_the_World_1fd01f.txt, Apsley_Cherry-Garrard_ff4e03.txt, Apsley_Cherry-Garrard_ff4e03.txt (+2) | Apsley_Cherry-Garrard_ff4e03.txt, The_Worst_Journey_in_the_World_1fd01f.txt | ✅ |
| q191 | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt, James_Fieser_e21429.txt (+2) | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | ✅ |
| q192 | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, The_Prince_and_Me_253bec.txt, Lee_Seung-gi_1ed0e1.txt (+2) | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, Lee_Seung-gi_1ed0e1.txt | ✅ |
| q193 | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+2) | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt | ✅ |
| q194 | Thomas_Doherty__actor__7caf88.txt, The_Lodge__TV_series__b905ce.txt, Life_with_Derek_547274.txt (+2) | The_Lodge__TV_series__b905ce.txt, Thomas_Doherty__actor__7caf88.txt | ✅ |
| q195 | Chelsea_Peretti_198d48.txt, Brooklyn_Nine-Nine_eec7c9.txt, Larry_Drake_85028c.txt (+2) | Brooklyn_Nine-Nine_eec7c9.txt, Chelsea_Peretti_198d48.txt | ✅ |
| q196 | Mark_Gaudet_0f4aa3.txt, Jan_Axel_Blomberg_a42602.txt, Mark_Lindsay_c8bd25.txt (+2) | Jan_Axel_Blomberg_a42602.txt, Mark_Gaudet_0f4aa3.txt | ✅ |
| q197 | Miami_Canal_161997.txt, Dundee_Canal_05ea0b.txt, Richard_Hornsby__amp__Sons_fb744c.txt (+2) | Dundee_Canal_05ea0b.txt, Miami_Canal_161997.txt | ✅ |
| q198 | Tron_8f60c9.txt, The_Million_Dollar_Duck_d1d45c.txt, Gryphon__film__f811a3.txt (+2) | The_Million_Dollar_Duck_d1d45c.txt, Tron_8f60c9.txt | ✅ |
| q199 | Beauty_and_the_Beast__franchise__7a780a.txt, Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt (+2) | Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt | ✅ |
| q200 | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt, Peter_Schmeichel_0b144a.txt (+2) | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt | ✅ |

### Config: top10 (top_k=10)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt, Eski_Imaret_Mosque_b4d62c.txt (+7) | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt | ✅ |
| q2 | Random_House_Tower_2b6a2a.txt, 888_7th_Avenue_f091f0.txt, 888_7th_Avenue_f091f0.txt (+7) | 888_7th_Avenue_f091f0.txt, Random_House_Tower_2b6a2a.txt | ✅ |
| q3 | Alex_Ferguson_6962e9.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt, Peter_Schmeichel_0b144a.txt (+7) | 1995_96_Manchester_United_F.C._season_3f9756.txt, Alex_Ferguson_6962e9.txt | ✅ |
| q4 | Apple_Remote_fcaffa.txt, Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt (+7) | Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt | ✅ |
| q5 | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt, Charles_Nungesser_67ba12.txt (+7) | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt | ✅ |
| q6 | Henry_J._Kaiser_53c448.txt, Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt (+7) | Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt | ✅ |
| q7 | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt, Jean-Loup_Chrétien_1d1a07.txt (+7) | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt | ✅ |
| q8 | Freakonomics__film__249214.txt, In_the_Realm_of_the_Hackers_367f43.txt, Connections__TV_series__deb419.txt (+7) | In_the_Realm_of_the_Hackers_367f43.txt, Freakonomics__film__249214.txt | ✅ |
| q9 | Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt (+7) | Socialist_Revolutionary_Party_68737a.txt, Russian_Civil_War_5a6752.txt | ✅ |
| q10 | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt, Ogallala__Nebraska_3578d6.txt (+7) | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt | ✅ |
| q11 | Giuseppe_Arimondi_0fc7eb.txt, Battle_of_Adwa_1fa890.txt, Addis_Ababa_6c835e.txt (+7) | Battle_of_Adwa_1fa890.txt, Giuseppe_Arimondi_0fc7eb.txt | ✅ |
| q12 | Dirleton_Castle_a042b8.txt, Yellowcraigs_0c745c.txt, Kingdom_of_Northumbria_4036b8.txt (+7) | Yellowcraigs_0c745c.txt, Dirleton_Castle_a042b8.txt | ✅ |
| q13 | English_Electric_Canberra_becfbe.txt, Avro_Vulcan_7aa981.txt, English_Electric_Canberra_becfbe.txt (+7) | English_Electric_Canberra_becfbe.txt, No._2_Squadron_RAAF_a75482.txt | ✅ |
| q14 | Euromarché_085339.txt, Carrefour_a2a75b.txt, Maxeda_f76dfe.txt (+7) | Euromarché_085339.txt, Carrefour_a2a75b.txt | ✅ |
| q15 | Delirium__Ellie_Goulding_album__5bb0cb.txt, On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Erika_Jayne_60e347.txt (+7) | On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Delirium__Ellie_Goulding_album__5bb0cb.txt | ✅ |
| q16 | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt, The_Legend_of_Korra_86cdc2.txt (+7) | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt | ✅ |
| q17 | Oranjegekte_8248a5.txt, Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt (+7) | Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt | ✅ |
| q18 | Tromeo_and_Juliet_86fab3.txt, James_Gunn_9a06f8.txt, Romeo_87e7c1.txt (+7) | James_Gunn_9a06f8.txt, Tromeo_and_Juliet_86fab3.txt | ✅ |
| q19 | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt, Bob_Seger_69d05d.txt (+7) | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt | ✅ |
| q20 | Rostker_v._Goldberg_b61238.txt, Conscription_in_the_United_States_e920fa.txt, Conscription_in_the_United_States_e920fa.txt (+7) | Conscription_in_the_United_States_e920fa.txt, Rostker_v._Goldberg_b61238.txt | ✅ |
| q21 | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt, Orange_Julius_e106c8.txt (+7) | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt | ✅ |
| q22 | Their_Lives_eed122.txt, Monica_Lewinsky_7bb6c2.txt, Nancy_Soderberg_a37451.txt (+7) | Monica_Lewinsky_7bb6c2.txt, Their_Lives_eed122.txt | ✅ |
| q23 | Teide_National_Park_aaf674.txt, Garajonay_National_Park_97f362.txt, Hatton_Castle__Angus_eb96ea.txt (+7) | Garajonay_National_Park_97f362.txt, Teide_National_Park_aaf674.txt | ✅ |
| q24 | Andrew_Jaspan_f6dc15.txt, Andrew_Jaspan_f6dc15.txt, The_Conversation__website__724191.txt (+7) | The_Conversation__website__724191.txt, Andrew_Jaspan_f6dc15.txt | ✅ |
| q25 | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt, The_Conversation__website__724191.txt (+7) | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt | ✅ |
| q26 | Tysons_Galleria_ead975.txt, Oldham_County__Kentucky_34ed53.txt, Ogallala__Nebraska_3578d6.txt (+7) | Tysons_Galleria_ead975.txt, McLean__Virginia_45ec68.txt | ✅ |
| q27 | My_Eyes__Blake_Shelton_song__fa840b.txt, Based_on_a_True_Story..._42790b.txt, The_Best_of_LeAnn_Rimes_722f10.txt (+7) | Based_on_a_True_Story..._42790b.txt, My_Eyes__Blake_Shelton_song__fa840b.txt | ✅ |
| q28 | Caroline_Carver__actress__7577ed.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Carol__film__cbfb1f.txt (+7) | The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Caroline_Carver__actress__7577ed.txt | ✅ |
| q29 | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt, Wilton_Mall_66971e.txt (+7) | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt | ✅ |
| q30 | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt, Jessica_Rothe_59deb4.txt (+7) | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt | ✅ |
| q31 | Mummulgum_1292e7.txt, Casino__New_South_Wales_8c85c5.txt, Ogallala__Nebraska_3578d6.txt (+7) | Casino__New_South_Wales_8c85c5.txt, Mummulgum_1292e7.txt | ✅ |
| q32 | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt, Sacred_Planet_26fd7b.txt (+7) | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt | ✅ |
| q33 | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt, James_Gunn_9a06f8.txt (+7) | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt | ✅ |
| q34 | Roberta_Vinci_f714c8.txt, Jorge_Lozano_1baa7a.txt, Jorge_Lozano_1baa7a.txt (+7) | Jorge_Lozano_1baa7a.txt, Roberta_Vinci_f714c8.txt | ✅ |
| q35 | Marco_Da_Silva__dancer__777e91.txt, Erika_Jayne_60e347.txt, Cressida_Bonas_5c3c57.txt (+7) | Erika_Jayne_60e347.txt, Marco_Da_Silva__dancer__777e91.txt | ✅ |
| q36 | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt (+7) | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt | ✅ |
| q37 | Kingdom_of_the_Isles_9a036f.txt, Kingdom_of_the_Isles_9a036f.txt, Aonghus_Mór_b5e643.txt (+7) | Aonghus_Mór_b5e643.txt, Kingdom_of_the_Isles_9a036f.txt | ✅ |
| q38 | Bruce_Spizer_c78d7e.txt, Bob_Seger_69d05d.txt, The_Beatles_c9e770.txt (+7) | The_Beatles_c9e770.txt, Bruce_Spizer_c78d7e.txt | ✅ |
| q39 | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt, Wayne_Garland_02f688.txt (+7) | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt | ✅ |
| q40 | Argand_lamp_563fb5.txt, Lewis_lamp_ddcc57.txt, Lewis_lamp_ddcc57.txt (+7) | Lewis_lamp_ddcc57.txt, Argand_lamp_563fb5.txt | ✅ |
| q41 | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt, Kathy_Sullivan__Australian_politician__a2272b.txt (+7) | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt | ✅ |
| q42 | Bishop_Carroll_Catholic_High_School_5aa312.txt, Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+7) | Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt | ✅ |
| q43 | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt, United_States_presidential_election__2016_363dfd.txt (+7) | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt | ✅ |
| q44 | Southaven__Mississippi_7cc376.txt, Memphis_Hustle_71bf0b.txt, Ogallala__Nebraska_3578d6.txt (+7) | Memphis_Hustle_71bf0b.txt, Southaven__Mississippi_7cc376.txt | ✅ |
| q45 | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt, Something_There_380707.txt (+7) | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt | ✅ |
| q46 | Albertina_eec665.txt, Hanna_Varis_e45643.txt, Hanna_Varis_e45643.txt (+7) | Albertina_eec665.txt, Hanna_Varis_e45643.txt | ✅ |
| q47 | Hatton_Hill_027377.txt, Hatton_Castle__Angus_eb96ea.txt, Hatton_Castle__Angus_eb96ea.txt (+7) | Hatton_Castle__Angus_eb96ea.txt, Hatton_Hill_027377.txt | ✅ |
| q48 | The_Legend_of_Korra_86cdc2.txt, Kuvira_4fdaa0.txt, Gargoyles__TV_series__acb424.txt (+7) | Kuvira_4fdaa0.txt, The_Legend_of_Korra_86cdc2.txt | ✅ |
| q49 | The_Five_Obstructions_18f018.txt, The_Importance_of_Being_Icelandic_086a7b.txt, The_Importance_of_Being_Icelandic_086a7b.txt (+7) | The_Importance_of_Being_Icelandic_086a7b.txt, The_Five_Obstructions_18f018.txt | ✅ |
| q50 | Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt, Will__amp__Grace_31b102.txt, The_Legend_of_Korra_86cdc2.txt (+7) | Will__amp__Grace_31b102.txt, Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt | ✅ |
| q51 | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt, Kohlberg_Kravis_Roberts_728df7.txt (+7) | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt | ✅ |
| q52 | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt, Wendell_Berry_4313ab.txt (+7) | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt | ✅ |
| q53 | 1920__film__db5e98.txt, Soha_Ali_Khan_b0280e.txt, 1920__film_series__78d752.txt (+7) | 1920__film_series__78d752.txt, 1920__film__db5e98.txt | ✅ |
| q54 | 71st_Golden_Globe_Awards_8dbed6.txt, 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt (+7) | 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt | ✅ |
| q55 | Charles_Hastings_Judd_409cf1.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt, Kalākaua_055843.txt (+7) | Charles_Hastings_Judd_409cf1.txt, Kalākaua_055843.txt | ✅ |
| q56 | Armie_Hammer_7c1778.txt, The_Polar_Bears_b06086.txt, Eddie_Izzard_fb2be0.txt (+7) | The_Polar_Bears_b06086.txt, Armie_Hammer_7c1778.txt | ✅ |
| q57 | 712_Fifth_Avenue_4807f9.txt, Manhattan_Life_Insurance_Building_fb6468.txt, Random_House_Tower_2b6a2a.txt (+7) | Manhattan_Life_Insurance_Building_fb6468.txt, 712_Fifth_Avenue_4807f9.txt | ✅ |
| q58 | Tenerife_c81266.txt, Gerald_Reive_c5cd23.txt, Samoa_086113.txt (+7) | Samoa_086113.txt, Gerald_Reive_c5cd23.txt | ✅ |
| q59 | Tecumseh_bcf68b.txt, Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt (+7) | Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt | ✅ |
| q60 | Samuel_Sim_07d71f.txt, Tromeo_and_Juliet_86fab3.txt, Bedknobs_and_Broomsticks_090f32.txt (+7) | Samuel_Sim_07d71f.txt, Awake__film__360ee6.txt | ✅ |
| q61 | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt, Øresund_Region_549ba5.txt (+7) | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt | ✅ |
| q62 | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt, Larry_Drake_85028c.txt (+7) | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt | ✅ |
| q63 | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt, Herbert_Akroyd_Stuart_0222bb.txt (+7) | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt | ✅ |
| q64 | United_States_v._Paramount_Pictures__Inc._4aa665.txt, Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+7) | Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | ✅ |
| q65 | Children_s_Mercy_Park_c9283f.txt, Arrowhead_Stadium_eae21e.txt, CommunityAmerica_Ballpark_1e3f7d.txt (+7) | Children_s_Mercy_Park_c9283f.txt, CommunityAmerica_Ballpark_1e3f7d.txt | ✅ |
| q66 | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt, The_Informant__376ad6.txt (+7) | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt | ✅ |
| q67 | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt, Eucryphia_3801b3.txt (+7) | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt | ✅ |
| q68 | Something_There_380707.txt, Paige_O_Hara_4492a0.txt, Something_There_380707.txt (+7) | Paige_O_Hara_4492a0.txt, Something_There_380707.txt | ✅ |
| q69 | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt, Laleli_Mosque_c7818f.txt (+7) | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt | ✅ |
| q70 | Opry_Mills_53c375.txt, Music_City_Queen_82e4a1.txt, Southaven__Mississippi_7cc376.txt (+7) | Music_City_Queen_82e4a1.txt, Opry_Mills_53c375.txt | ✅ |
| q71 | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt, Opry_Mills_53c375.txt (+7) | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt | ✅ |
| q72 | James_Fieser_e21429.txt, Berea_College_88c6f9.txt, Azusa_Pacific_University_fcee08.txt (+7) | James_Fieser_e21429.txt, Berea_College_88c6f9.txt | ✅ |
| q73 | James_Burke__science_historian__6fe4bf.txt, Connections__TV_series__deb419.txt, Connections__TV_series__deb419.txt (+7) | Connections__TV_series__deb419.txt, James_Burke__science_historian__6fe4bf.txt | ✅ |
| q74 | Romeo_87e7c1.txt, Benvolio_848ddf.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt (+7) | Romeo_87e7c1.txt, Benvolio_848ddf.txt | ✅ |
| q75 | Addis_Ababa_6c835e.txt, National_Archives_and_Library_of_Ethiopia_6068a1.txt, Ogallala__Nebraska_3578d6.txt (+7) | National_Archives_and_Library_of_Ethiopia_6068a1.txt, Addis_Ababa_6c835e.txt | ✅ |
| q76 | Night_Ferry__composition__9f4c6a.txt, Toshi_Ichiyanagi_e50d8d.txt, Symphony_Center_f29c57.txt (+7) | Night_Ferry__composition__9f4c6a.txt, Symphony_Center_f29c57.txt | ✅ |
| q77 | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt, Laura_Osnes_4365e8.txt (+7) | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt | ✅ |
| q78 | Eucryphia_3801b3.txt, Lepidozamia_4d1d3c.txt, Elatostema_fff908.txt (+7) | Lepidozamia_4d1d3c.txt, Eucryphia_3801b3.txt | ✅ |
| q79 | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt, Samoa_086113.txt (+7) | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt | ✅ |
| q80 | Kris_Marshall_97e2d7.txt, Death_in_Paradise__TV_series__8a650d.txt, Eddie_Izzard_fb2be0.txt (+7) | Death_in_Paradise__TV_series__8a650d.txt, Kris_Marshall_97e2d7.txt | ✅ |
| q81 | EgyptAir_Flight_990_dfd74a.txt, Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt (+7) | Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt | ✅ |
| q82 | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt, Oz_the_Great_and_Powerful_510b50.txt (+7) | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt | ✅ |
| q83 | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt, Jacques_Sernas_2c77cc.txt (+7) | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt | ✅ |
| q84 | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt, Samoa_086113.txt (+7) | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt | ✅ |
| q85 | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt, Second_Anglo-Afghan_War_adb0b2.txt (+7) | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt | ✅ |
| q86 | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt, Bolton_6ce6c9.txt (+7) | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt | ✅ |
| q87 | Hot_air_engine_2ed8e1.txt, Herbert_Akroyd_Stuart_0222bb.txt, George_Cayley_2c8397.txt (+7) | George_Cayley_2c8397.txt, Hot_air_engine_2ed8e1.txt | ✅ |
| q88 | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt, Beauty_and_the_Beast__1991_film__d38192.txt (+7) | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt | ✅ |
| q89 | Northumbrian_dialect_7334ca.txt, Kingdom_of_Northumbria_4036b8.txt, Kingdom_of_the_Isles_9a036f.txt (+7) | Kingdom_of_Northumbria_4036b8.txt, Northumbrian_dialect_7334ca.txt | ✅ |
| q90 | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt, Southaven__Mississippi_7cc376.txt (+7) | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt | ✅ |
| q91 | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt (+7) | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt | ✅ |
| q92 | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+7) | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt | ✅ |
| q93 | Oedipus_Rex_d47dfb.txt, Dostoevsky_and_Parricide_f04c2c.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+7) | Dostoevsky_and_Parricide_f04c2c.txt, Oedipus_Rex_d47dfb.txt | ✅ |
| q94 | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt, Chelsea_Peretti_198d48.txt (+7) | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt | ✅ |
| q95 | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt, Kunming_08b9a1.txt (+7) | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt | ✅ |
| q96 | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+7) | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt | ✅ |
| q97 | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt, Adnan_Akmal_e863b9.txt (+7) | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt | ✅ |
| q98 | Arrowhead_Stadium_eae21e.txt, Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt (+7) | Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt | ✅ |
| q99 | Happy_Death_Day_793b37.txt, Jessica_Rothe_59deb4.txt, The_Bye_Bye_Man_906e6d.txt (+7) | Jessica_Rothe_59deb4.txt, Happy_Death_Day_793b37.txt | ✅ |
| q100 | Garden_Island_Naval_Chapel_6a3c4e.txt, Royal_Australian_Navy_5e1d16.txt, Samoa_086113.txt (+7) | Royal_Australian_Navy_5e1d16.txt, Garden_Island_Naval_Chapel_6a3c4e.txt | ✅ |
| q101 | The_Informant__376ad6.txt, Mark_Whitacre_ccc607.txt, Awake__film__360ee6.txt (+7) | Mark_Whitacre_ccc607.txt, The_Informant__376ad6.txt | ✅ |
| q102 | Current_Mood_0317b8.txt, Small_Town_Boy__song__d2fddb.txt, Based_on_a_True_Story..._42790b.txt (+7) | Small_Town_Boy__song__d2fddb.txt, Current_Mood_0317b8.txt | ✅ |
| q103 | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt, Impresario_5cc7af.txt (+7) | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt | ✅ |
| q104 | Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt, Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt (+7) | Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt | ✅ |
| q105 | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt, Tim_Hecker_80429a.txt (+7) | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt | ✅ |
| q106 | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt, Adrian_Peterson_1ca55b.txt (+7) | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt | ✅ |
| q107 | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt, Goetheanum_5744ff.txt (+7) | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt | ✅ |
| q108 | The_Ganymede_Takeover_5feff5.txt, The_Man_in_the_High_Castle_1c064a.txt, Wendell_Berry_4313ab.txt (+7) | The_Man_in_the_High_Castle_1c064a.txt, The_Ganymede_Takeover_5feff5.txt | ✅ |
| q109 | Curt_Menefee_0ab3d9.txt, Michael_Strahan_1fa88f.txt, Michael_Strahan_1fa88f.txt (+7) | Michael_Strahan_1fa88f.txt, Curt_Menefee_0ab3d9.txt | ✅ |
| q110 | Summer_of_the_Monkeys_0b84ea.txt, William_Allen_White_d1418b.txt, T._R._M._Howard_0bb121.txt (+7) | William_Allen_White_d1418b.txt, Summer_of_the_Monkeys_0b84ea.txt | ✅ |
| q111 | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt, Shoba_Chandrasekhar_5285b2.txt (+7) | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt | ✅ |
| q112 | Alice_Upside_Down_2fe728.txt, Lucy_Fry_218cb3.txt, Caroline_Carver__actress__7577ed.txt (+7) | Lucas_Grabeel_c00cb8.txt, Alice_Upside_Down_2fe728.txt | ✅ |
| q113 | Snowdrop__game_engine__750d41.txt, Tom_Clancy_s_The_Division_3e09f9.txt, Icehouse_pieces_5e75b5.txt (+7) | Tom_Clancy_s_The_Division_3e09f9.txt, Snowdrop__game_engine__750d41.txt | ✅ |
| q114 | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt, Tom_Clancy_s_The_Division_3e09f9.txt (+7) | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt | ✅ |
| q115 | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt, Jean-Loup_Chrétien_1d1a07.txt (+7) | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt | ✅ |
| q116 | Banshee_5e6ebd.txt, VMAQT-1_05fe33.txt, VMAQT-1_05fe33.txt (+7) | VMAQT-1_05fe33.txt, Banshee_5e6ebd.txt | ✅ |
| q117 | Barbara_Niven_fbf739.txt, Awake__film__360ee6.txt, Alice_Upside_Down_2fe728.txt (+7) | Dead_at_17_273b88.txt, Barbara_Niven_fbf739.txt | ✅ |
| q118 | Bart_the_Fink_77de9b.txt, Krusty_the_Clown_3e3656.txt, Cedric_the_Entertainer_e39d6e.txt (+7) | Krusty_the_Clown_3e3656.txt, Bart_the_Fink_77de9b.txt | ✅ |
| q119 | Viaport_Rotterdam_760a19.txt, Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt (+7) | Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt | ✅ |
| q120 | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt, Marco_Da_Silva__dancer__777e91.txt (+7) | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt | ✅ |
| q121 | Ambrose_Mendy_eddce7.txt, Chris_Eubank_Jr._82fe88.txt, Peter_Schmeichel_0b144a.txt (+7) | Chris_Eubank_Jr._82fe88.txt, Ambrose_Mendy_eddce7.txt | ✅ |
| q122 | Allen__amp__Company_Sun_Valley_Conference_083cbf.txt, Rupert_Murdoch_8801f1.txt, Joe_Scarborough_fde209.txt (+7) | Rupert_Murdoch_8801f1.txt, Allen__amp__Company_Sun_Valley_Conference_083cbf.txt | ✅ |
| q123 | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt, Raymond_Ochoa_da4d56.txt (+7) | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt | ✅ |
| q124 | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt, Phoenix_Television_29103f.txt (+7) | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt | ✅ |
| q125 | Patricia_Longo_b7fcef.txt, Graduados_7592c5.txt, Tenerife_c81266.txt (+7) | Graduados_7592c5.txt, Patricia_Longo_b7fcef.txt | ✅ |
| q126 | Ogallala_Aquifer_a2b49c.txt, Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt (+7) | Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt | ✅ |
| q127 | Blinding_Edge_Pictures_b8de5a.txt, Unbreakable__film__52d8de.txt, Tron_8f60c9.txt (+7) | Unbreakable__film__52d8de.txt, Blinding_Edge_Pictures_b8de5a.txt | ✅ |
| q128 | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt, The_Good_Dinosaur_170ac4.txt (+7) | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt | ✅ |
| q129 | BraveStarr_8d412d.txt, Celebrity_Home_Entertainment_c01bf6.txt, Gargoyles__TV_series__acb424.txt (+7) | Celebrity_Home_Entertainment_c01bf6.txt, BraveStarr_8d412d.txt | ✅ |
| q130 | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt, The_Informant__376ad6.txt (+7) | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt | ✅ |
| q131 | Lucy_Fry_218cb3.txt, Jessica_Rothe_59deb4.txt, Nina_Dobrev_05e14f.txt (+7) | Mr._Church_ce0d51.txt, Lucy_Fry_218cb3.txt | ✅ |
| q132 | Shoba_Chandrasekhar_5285b2.txt, Ithu_Engal_Neethi_ad89e5.txt, Soha_Ali_Khan_b0280e.txt (+7) | Ithu_Engal_Neethi_ad89e5.txt, Shoba_Chandrasekhar_5285b2.txt | ✅ |
| q133 | Official_Ireland_2ed543.txt, Catholic_Church_in_Ireland_77cac6.txt, Catholic_Church_in_Ireland_77cac6.txt (+7) | Catholic_Church_in_Ireland_77cac6.txt, Official_Ireland_2ed543.txt | ✅ |
| q134 | Bridge_to_Terabithia__1985_film__1aaa6c.txt, Bedknobs_and_Broomsticks_090f32.txt, Bridge_to_Terabithia__novel__21de92.txt (+7) | Bridge_to_Terabithia__novel__21de92.txt, Bridge_to_Terabithia__1985_film__1aaa6c.txt | ✅ |
| q135 | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt, Curt_Menefee_0ab3d9.txt (+7) | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt | ✅ |
| q136 | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt, Kasper_Schmeichel_c9da28.txt (+7) | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt | ✅ |
| q137 | Atari_Assembler_Editor_a5f9fc.txt, Shepardson_Microsystems_0fa820.txt, Snowdrop__game_engine__750d41.txt (+7) | Shepardson_Microsystems_0fa820.txt, Atari_Assembler_Editor_a5f9fc.txt | ✅ |
| q138 | His_Band_and_the_Street_Choir_f5b88a.txt, I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt (+7) | I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt | ✅ |
| q139 | Aldosterone_b32476.txt, Aldosterone_b32476.txt, Angiotensin_3f2772.txt (+7) | Angiotensin_3f2772.txt, Aldosterone_b32476.txt | ✅ |
| q140 | Nancy_Soderberg_a37451.txt, United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt (+7) | United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt | ✅ |
| q141 | Mississippi_University_for_Women_v._Hogan_f450c0.txt, Berghuis_v._Thompkins_c16323.txt, Reynolds_v._Sims_31c83e.txt (+7) | Berghuis_v._Thompkins_c16323.txt, Mississippi_University_for_Women_v._Hogan_f450c0.txt | ✅ |
| q142 | AFC_North_8f76d9.txt, AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt (+7) | AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt | ✅ |
| q143 | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt (+7) | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt | ✅ |
| q144 | Honda_Ballade_4d8914.txt, Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt (+7) | Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt | ✅ |
| q145 | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt, The_Monster__song__886456.txt (+7) | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt | ✅ |
| q146 | Backstage__magazine__258d77.txt, Celebrity_Home_Entertainment_c01bf6.txt, Christopher_Oscar_Peña_3991f4.txt (+7) | Backstage__magazine__258d77.txt, Christopher_Oscar_Peña_3991f4.txt | ✅ |
| q147 | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt, Music_for_Electric_Metronomes_0c93d9.txt (+7) | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt | ✅ |
| q148 | Ego_the_Living_Planet_f6f847.txt, Guardians_of_the_Galaxy_Vol._2_b6c488.txt, James_Gunn_9a06f8.txt (+7) | Guardians_of_the_Galaxy_Vol._2_b6c488.txt, Ego_the_Living_Planet_f6f847.txt | ✅ |
| q149 | Sponsorship_scandal_5e9ed3.txt, Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt (+7) | Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt | ✅ |
| q150 | Conrad_Anker_9d25ca.txt, George_Mallory_ccb70e.txt, Apsley_Cherry-Garrard_ff4e03.txt (+7) | George_Mallory_ccb70e.txt, Conrad_Anker_9d25ca.txt | ✅ |
| q151 | Joan_Crawford_5fcb43.txt, The_Duke_Steps_Out_7cc9f2.txt, Erika_Jayne_60e347.txt (+7) | The_Duke_Steps_Out_7cc9f2.txt, Joan_Crawford_5fcb43.txt | ✅ |
| q152 | Carol__film__cbfb1f.txt, Guild_of_Music_Supervisors_Awards_7984fd.txt, The_Muppet_Christmas_Carol_81e722.txt (+7) | Guild_of_Music_Supervisors_Awards_7984fd.txt, Carol__film__cbfb1f.txt | ✅ |
| q153 | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+7) | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt | ✅ |
| q154 | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt (+7) | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt | ✅ |
| q155 | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt, Goetheanum_5744ff.txt (+7) | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt | ✅ |
| q156 | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt, Berea_College_88c6f9.txt (+7) | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt | ✅ |
| q157 | Sprawl_trilogy_894d39.txt, Neuromancer_6029bf.txt, The_Man_in_the_High_Castle_1c064a.txt (+7) | Neuromancer_6029bf.txt, Sprawl_trilogy_894d39.txt | ✅ |
| q158 | Azad_Hind_Dal_a97cda.txt, Subhas_Chandra_Bose_4b75c6.txt, Second_Anglo-Afghan_War_adb0b2.txt (+7) | Subhas_Chandra_Bose_4b75c6.txt, Azad_Hind_Dal_a97cda.txt | ✅ |
| q159 | Eddie_Izzard_fb2be0.txt, Stripped__tour__91c664.txt, Chelsea_Peretti_198d48.txt (+7) | Stripped__tour__91c664.txt, Eddie_Izzard_fb2be0.txt | ✅ |
| q160 | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt, Soha_Ali_Khan_b0280e.txt (+7) | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt | ✅ |
| q161 | Ecballium_ad2282.txt, Elatostema_fff908.txt, Thalictrum_3222b5.txt (+7) | Elatostema_fff908.txt, Ecballium_ad2282.txt | ✅ |
| q162 | Polypodium_71abb6.txt, Aichryson_319a15.txt, Thalictrum_3222b5.txt (+7) | Polypodium_71abb6.txt, Aichryson_319a15.txt | ✅ |
| q163 | Adoption_2002_46d793.txt, Adoption_and_Safe_Families_Act_b4c27b.txt, Conscription_in_the_United_States_e920fa.txt (+7) | Adoption_and_Safe_Families_Act_b4c27b.txt, Adoption_2002_46d793.txt | ✅ |
| q164 | Crash_Pad_ed95da.txt, Nina_Dobrev_05e14f.txt, The_Lodge__TV_series__b905ce.txt (+7) | Nina_Dobrev_05e14f.txt, Crash_Pad_ed95da.txt | ✅ |
| q165 | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt, Their_Lives_eed122.txt (+7) | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt | ✅ |
| q166 | Parks_and_Recreation_051519.txt, Ms._Knope_Goes_to_Washington_799ad4.txt, Will__amp__Grace_31b102.txt (+7) | Ms._Knope_Goes_to_Washington_799ad4.txt, Parks_and_Recreation_051519.txt | ✅ |
| q167 | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt, Kasper_Schmeichel_c9da28.txt (+7) | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt | ✅ |
| q168 | Can_t_Fight_the_Moonlight_caf06b.txt, The_Best_of_LeAnn_Rimes_722f10.txt, Tron_8f60c9.txt (+7) | The_Best_of_LeAnn_Rimes_722f10.txt, Can_t_Fight_the_Moonlight_caf06b.txt | ✅ |
| q169 | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt, Lee_Seung-gi_1ed0e1.txt (+7) | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt | ✅ |
| q170 | Unbelievable__The_Notorious_B.I.G._song__91da75.txt, Ready_to_Die_c706b2.txt, Ready_to_Die_c706b2.txt (+7) | Ready_to_Die_c706b2.txt, Unbelievable__The_Notorious_B.I.G._song__91da75.txt | ✅ |
| q171 | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt, Cedric_the_Entertainer_e39d6e.txt (+7) | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt | ✅ |
| q172 | Violin_Sonata_No._5__Beethoven__8f6942.txt, Symphony_No._7__Beethoven__9c3b01.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+7) | Symphony_No._7__Beethoven__9c3b01.txt, Violin_Sonata_No._5__Beethoven__8f6942.txt | ✅ |
| q173 | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt, Mondelez_International_9a57f2.txt (+7) | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt | ✅ |
| q174 | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt, Cedric_the_Entertainer_e39d6e.txt (+7) | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt | ✅ |
| q175 | Saab_36_108e14.txt, Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt (+7) | Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt | ✅ |
| q176 | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt, Pasek_and_Paul_aa5312.txt (+7) | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt | ✅ |
| q177 | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt, North_American_T-6_Texan_76936b.txt (+7) | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt | ✅ |
| q178 | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt, T._R._M._Howard_0bb121.txt (+7) | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt | ✅ |
| q179 | Lucas_Carvalho_940f12.txt, 4___400_metres_relay_8bba84.txt, 4___400_metres_relay_8bba84.txt (+7) | 4___400_metres_relay_8bba84.txt, Lucas_Carvalho_940f12.txt | ✅ |
| q180 | Watercliffe_Meadow_Community_Primary_School_1e4bca.txt, Political_correctness_f6d0c6.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+7) | Political_correctness_f6d0c6.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt | ✅ |
| q181 | Parodia_a9ceb8.txt, Thalictrum_3222b5.txt, Polypodium_71abb6.txt (+7) | Thalictrum_3222b5.txt, Parodia_a9ceb8.txt | ✅ |
| q182 | The_Simpsons__season_23__36cc7b.txt, At_Long_Last_Leave_8151f2.txt, Bart_the_Fink_77de9b.txt (+7) | At_Long_Last_Leave_8151f2.txt, The_Simpsons__season_23__36cc7b.txt | ✅ |
| q183 | BJ_s_Wholesale_Club_f114e3.txt, US_Vision_f98d42.txt, US_Vision_f98d42.txt (+7) | US_Vision_f98d42.txt, BJ_s_Wholesale_Club_f114e3.txt | ✅ |
| q184 | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt, Lochmere_Archeological_District_2ae282.txt (+7) | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt | ✅ |
| q185 | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt, Southaven__Mississippi_7cc376.txt (+7) | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt | ✅ |
| q186 | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt, Samantha_Cristoforetti_fc1bcd.txt (+7) | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt | ✅ |
| q187 | Flynn_Rider_37adf2.txt, Zachary_Levi_9c18c7.txt, Lucas_Grabeel_c00cb8.txt (+7) | Zachary_Levi_9c18c7.txt, Flynn_Rider_37adf2.txt | ✅ |
| q188 | Vacation_with_Derek_da8ab6.txt, Life_with_Derek_547274.txt, Cressida_Bonas_5c3c57.txt (+7) | Life_with_Derek_547274.txt, Vacation_with_Derek_da8ab6.txt | ✅ |
| q189 | Jin_Jing_fa2e08.txt, Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt (+7) | Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt | ✅ |
| q190 | The_Worst_Journey_in_the_World_1fd01f.txt, Apsley_Cherry-Garrard_ff4e03.txt, Apsley_Cherry-Garrard_ff4e03.txt (+7) | Apsley_Cherry-Garrard_ff4e03.txt, The_Worst_Journey_in_the_World_1fd01f.txt | ✅ |
| q191 | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt, James_Fieser_e21429.txt (+7) | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | ✅ |
| q192 | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, The_Prince_and_Me_253bec.txt, Lee_Seung-gi_1ed0e1.txt (+7) | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, Lee_Seung-gi_1ed0e1.txt | ✅ |
| q193 | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+7) | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt | ✅ |
| q194 | Thomas_Doherty__actor__7caf88.txt, The_Lodge__TV_series__b905ce.txt, Life_with_Derek_547274.txt (+7) | The_Lodge__TV_series__b905ce.txt, Thomas_Doherty__actor__7caf88.txt | ✅ |
| q195 | Chelsea_Peretti_198d48.txt, Brooklyn_Nine-Nine_eec7c9.txt, Larry_Drake_85028c.txt (+7) | Brooklyn_Nine-Nine_eec7c9.txt, Chelsea_Peretti_198d48.txt | ✅ |
| q196 | Mark_Gaudet_0f4aa3.txt, Jan_Axel_Blomberg_a42602.txt, Mark_Lindsay_c8bd25.txt (+7) | Jan_Axel_Blomberg_a42602.txt, Mark_Gaudet_0f4aa3.txt | ✅ |
| q197 | Miami_Canal_161997.txt, Dundee_Canal_05ea0b.txt, Richard_Hornsby__amp__Sons_fb744c.txt (+7) | Dundee_Canal_05ea0b.txt, Miami_Canal_161997.txt | ✅ |
| q198 | Tron_8f60c9.txt, The_Million_Dollar_Duck_d1d45c.txt, Gryphon__film__f811a3.txt (+7) | The_Million_Dollar_Duck_d1d45c.txt, Tron_8f60c9.txt | ✅ |
| q199 | Beauty_and_the_Beast__franchise__7a780a.txt, Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt (+7) | Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt | ✅ |
| q200 | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt, Peter_Schmeichel_0b144a.txt (+7) | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt | ✅ |

### Config: top20 (top_k=20)

| Query | Retrieved | Relevant | Hit? |
|-------|-----------|----------|------|
| q1 | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt, Eski_Imaret_Mosque_b4d62c.txt (+17) | Laleli_Mosque_c7818f.txt, Esma_Sultan_Mansion_fb6370.txt | ✅ |
| q2 | Random_House_Tower_2b6a2a.txt, 888_7th_Avenue_f091f0.txt, 888_7th_Avenue_f091f0.txt (+17) | 888_7th_Avenue_f091f0.txt, Random_House_Tower_2b6a2a.txt | ✅ |
| q3 | Alex_Ferguson_6962e9.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt, Peter_Schmeichel_0b144a.txt (+17) | 1995_96_Manchester_United_F.C._season_3f9756.txt, Alex_Ferguson_6962e9.txt | ✅ |
| q4 | Apple_Remote_fcaffa.txt, Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt (+17) | Apple_Remote_fcaffa.txt, Front_Row__software__ec6691.txt | ✅ |
| q5 | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt, Charles_Nungesser_67ba12.txt (+17) | Kasper_Schmeichel_c9da28.txt, Peter_Schmeichel_0b144a.txt | ✅ |
| q6 | Henry_J._Kaiser_53c448.txt, Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt (+17) | Kaiser_Ventures_93a330.txt, Henry_J._Kaiser_53c448.txt | ✅ |
| q7 | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt, Jean-Loup_Chrétien_1d1a07.txt (+17) | L_Oiseau_Blanc_12450d.txt, Charles_Nungesser_67ba12.txt | ✅ |
| q8 | Freakonomics__film__249214.txt, In_the_Realm_of_the_Hackers_367f43.txt, Connections__TV_series__deb419.txt (+17) | In_the_Realm_of_the_Hackers_367f43.txt, Freakonomics__film__249214.txt | ✅ |
| q9 | Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt, Russian_Civil_War_5a6752.txt (+17) | Socialist_Revolutionary_Party_68737a.txt, Russian_Civil_War_5a6752.txt | ✅ |
| q10 | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt, Ogallala__Nebraska_3578d6.txt (+17) | Gerald_R._Ford_International_Airport_aa3aad.txt, Elko_Regional_Airport_b8dabb.txt | ✅ |
| q11 | Giuseppe_Arimondi_0fc7eb.txt, Battle_of_Adwa_1fa890.txt, Addis_Ababa_6c835e.txt (+17) | Battle_of_Adwa_1fa890.txt, Giuseppe_Arimondi_0fc7eb.txt | ✅ |
| q12 | Dirleton_Castle_a042b8.txt, Yellowcraigs_0c745c.txt, Kingdom_of_Northumbria_4036b8.txt (+17) | Yellowcraigs_0c745c.txt, Dirleton_Castle_a042b8.txt | ✅ |
| q13 | English_Electric_Canberra_becfbe.txt, Avro_Vulcan_7aa981.txt, English_Electric_Canberra_becfbe.txt (+17) | English_Electric_Canberra_becfbe.txt, No._2_Squadron_RAAF_a75482.txt | ✅ |
| q14 | Euromarché_085339.txt, Carrefour_a2a75b.txt, Maxeda_f76dfe.txt (+17) | Euromarché_085339.txt, Carrefour_a2a75b.txt | ✅ |
| q15 | Delirium__Ellie_Goulding_album__5bb0cb.txt, On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Erika_Jayne_60e347.txt (+17) | On_My_Mind__Ellie_Goulding_song__7beb6f.txt, Delirium__Ellie_Goulding_album__5bb0cb.txt | ✅ |
| q16 | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt, The_Legend_of_Korra_86cdc2.txt (+17) | Teen_Titans_Go___TV_series__f1f123.txt, Tara_Strong_a3946c.txt | ✅ |
| q17 | Oranjegekte_8248a5.txt, Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt (+17) | Koningsdag_f66ba4.txt, Oranjegekte_8248a5.txt | ✅ |
| q18 | Tromeo_and_Juliet_86fab3.txt, James_Gunn_9a06f8.txt, Romeo_87e7c1.txt (+17) | James_Gunn_9a06f8.txt, Tromeo_and_Juliet_86fab3.txt | ✅ |
| q19 | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt, Bob_Seger_69d05d.txt (+17) | Bob_Seger_69d05d.txt, Against_the_Wind__album__92e597.txt | ✅ |
| q20 | Rostker_v._Goldberg_b61238.txt, Conscription_in_the_United_States_e920fa.txt, Conscription_in_the_United_States_e920fa.txt (+17) | Conscription_in_the_United_States_e920fa.txt, Rostker_v._Goldberg_b61238.txt | ✅ |
| q21 | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt, Orange_Julius_e106c8.txt (+17) | Mondelez_International_9a57f2.txt, Handi-Snacks_b5858e.txt | ✅ |
| q22 | Their_Lives_eed122.txt, Monica_Lewinsky_7bb6c2.txt, Nancy_Soderberg_a37451.txt (+17) | Monica_Lewinsky_7bb6c2.txt, Their_Lives_eed122.txt | ✅ |
| q23 | Teide_National_Park_aaf674.txt, Garajonay_National_Park_97f362.txt, Hatton_Castle__Angus_eb96ea.txt (+17) | Garajonay_National_Park_97f362.txt, Teide_National_Park_aaf674.txt | ✅ |
| q24 | Andrew_Jaspan_f6dc15.txt, Andrew_Jaspan_f6dc15.txt, The_Conversation__website__724191.txt (+17) | The_Conversation__website__724191.txt, Andrew_Jaspan_f6dc15.txt | ✅ |
| q25 | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt, The_Conversation__website__724191.txt (+17) | London_International_Documentary_Festival_ee94b8.txt, London_Review_of_Books_e260ff.txt | ✅ |
| q26 | Tysons_Galleria_ead975.txt, Oldham_County__Kentucky_34ed53.txt, Ogallala__Nebraska_3578d6.txt (+17) | Tysons_Galleria_ead975.txt, McLean__Virginia_45ec68.txt | ✅ |
| q27 | My_Eyes__Blake_Shelton_song__fa840b.txt, Based_on_a_True_Story..._42790b.txt, The_Best_of_LeAnn_Rimes_722f10.txt (+17) | Based_on_a_True_Story..._42790b.txt, My_Eyes__Blake_Shelton_song__fa840b.txt | ✅ |
| q28 | Caroline_Carver__actress__7577ed.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Carol__film__cbfb1f.txt (+17) | The_Magical_Legend_of_the_Leprechauns_0d732f.txt, Caroline_Carver__actress__7577ed.txt | ✅ |
| q29 | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt, Wilton_Mall_66971e.txt (+17) | Maxeda_f76dfe.txt, Kohlberg_Kravis_Roberts_728df7.txt | ✅ |
| q30 | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt, Jessica_Rothe_59deb4.txt (+17) | Cressida_Bonas_5c3c57.txt, The_Bye_Bye_Man_906e6d.txt | ✅ |
| q31 | Mummulgum_1292e7.txt, Casino__New_South_Wales_8c85c5.txt, Ogallala__Nebraska_3578d6.txt (+17) | Casino__New_South_Wales_8c85c5.txt, Mummulgum_1292e7.txt | ✅ |
| q32 | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt, Sacred_Planet_26fd7b.txt (+17) | LaLee_s_Kin__The_Legacy_of_Cotton_6d9247.txt, Gimme_Shelter__1970_film__5d2f2a.txt | ✅ |
| q33 | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt, James_Gunn_9a06f8.txt (+17) | Harsh_Times_d738dc.txt, David_Ayer_4a06f2.txt | ✅ |
| q34 | Roberta_Vinci_f714c8.txt, Jorge_Lozano_1baa7a.txt, Jorge_Lozano_1baa7a.txt (+17) | Jorge_Lozano_1baa7a.txt, Roberta_Vinci_f714c8.txt | ✅ |
| q35 | Marco_Da_Silva__dancer__777e91.txt, Erika_Jayne_60e347.txt, Cressida_Bonas_5c3c57.txt (+17) | Erika_Jayne_60e347.txt, Marco_Da_Silva__dancer__777e91.txt | ✅ |
| q36 | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt (+17) | Reading_Post_22f3d7.txt, Maiwand_Lion_299aff.txt | ✅ |
| q37 | Kingdom_of_the_Isles_9a036f.txt, Kingdom_of_the_Isles_9a036f.txt, Aonghus_Mór_b5e643.txt (+17) | Aonghus_Mór_b5e643.txt, Kingdom_of_the_Isles_9a036f.txt | ✅ |
| q38 | Bruce_Spizer_c78d7e.txt, Bob_Seger_69d05d.txt, The_Beatles_c9e770.txt (+17) | The_Beatles_c9e770.txt, Bruce_Spizer_c78d7e.txt | ✅ |
| q39 | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt, Wayne_Garland_02f688.txt (+17) | Baltimore_Orioles_cf318a.txt, Wayne_Garland_02f688.txt | ✅ |
| q40 | Argand_lamp_563fb5.txt, Lewis_lamp_ddcc57.txt, Lewis_lamp_ddcc57.txt (+17) | Lewis_lamp_ddcc57.txt, Argand_lamp_563fb5.txt | ✅ |
| q41 | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt, Kathy_Sullivan__Australian_politician__a2272b.txt (+17) | Kathy_Sullivan__Australian_politician__a2272b.txt, Bronwyn_Bishop_60d0f7.txt | ✅ |
| q42 | Bishop_Carroll_Catholic_High_School_5aa312.txt, Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+17) | Kapaun_Mt._Carmel_Catholic_High_School_bf8c9d.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt | ✅ |
| q43 | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt, United_States_presidential_election__2016_363dfd.txt (+17) | Michigan_Democratic_primary__2016_7fa23f.txt, United_States_presidential_election__2016_363dfd.txt | ✅ |
| q44 | Southaven__Mississippi_7cc376.txt, Memphis_Hustle_71bf0b.txt, Ogallala__Nebraska_3578d6.txt (+17) | Memphis_Hustle_71bf0b.txt, Southaven__Mississippi_7cc376.txt | ✅ |
| q45 | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt, Something_There_380707.txt (+17) | Pasek_and_Paul_aa5312.txt, A_Christmas_Story__The_Musical_cc93e8.txt | ✅ |
| q46 | Albertina_eec665.txt, Hanna_Varis_e45643.txt, Hanna_Varis_e45643.txt (+17) | Albertina_eec665.txt, Hanna_Varis_e45643.txt | ✅ |
| q47 | Hatton_Hill_027377.txt, Hatton_Castle__Angus_eb96ea.txt, Hatton_Castle__Angus_eb96ea.txt (+17) | Hatton_Castle__Angus_eb96ea.txt, Hatton_Hill_027377.txt | ✅ |
| q48 | The_Legend_of_Korra_86cdc2.txt, Kuvira_4fdaa0.txt, Gargoyles__TV_series__acb424.txt (+17) | Kuvira_4fdaa0.txt, The_Legend_of_Korra_86cdc2.txt | ✅ |
| q49 | The_Five_Obstructions_18f018.txt, The_Importance_of_Being_Icelandic_086a7b.txt, The_Importance_of_Being_Icelandic_086a7b.txt (+17) | The_Importance_of_Being_Icelandic_086a7b.txt, The_Five_Obstructions_18f018.txt | ✅ |
| q50 | Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt, Will__amp__Grace_31b102.txt, The_Legend_of_Korra_86cdc2.txt (+17) | Will__amp__Grace_31b102.txt, Marry_Me_a_Little__Marry_Me_a_Little_More_8813fe.txt | ✅ |
| q51 | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt, Kohlberg_Kravis_Roberts_728df7.txt (+17) | Ravi_Sethi_7560a5.txt, Bell_Labs_726cbc.txt | ✅ |
| q52 | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt, Wendell_Berry_4313ab.txt (+17) | Dim_Gray_Bar_Press_809c57.txt, Wendell_Berry_4313ab.txt | ✅ |
| q53 | 1920__film__db5e98.txt, Soha_Ali_Khan_b0280e.txt, 1920__film_series__78d752.txt (+17) | 1920__film_series__78d752.txt, 1920__film__db5e98.txt | ✅ |
| q54 | 71st_Golden_Globe_Awards_8dbed6.txt, 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt (+17) | 71st_Golden_Globe_Awards_8dbed6.txt, Brooklyn_Nine-Nine_eec7c9.txt | ✅ |
| q55 | Charles_Hastings_Judd_409cf1.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt, Kalākaua_055843.txt (+17) | Charles_Hastings_Judd_409cf1.txt, Kalākaua_055843.txt | ✅ |
| q56 | Armie_Hammer_7c1778.txt, The_Polar_Bears_b06086.txt, Eddie_Izzard_fb2be0.txt (+17) | The_Polar_Bears_b06086.txt, Armie_Hammer_7c1778.txt | ✅ |
| q57 | 712_Fifth_Avenue_4807f9.txt, Manhattan_Life_Insurance_Building_fb6468.txt, Random_House_Tower_2b6a2a.txt (+17) | Manhattan_Life_Insurance_Building_fb6468.txt, 712_Fifth_Avenue_4807f9.txt | ✅ |
| q58 | Tenerife_c81266.txt, Gerald_Reive_c5cd23.txt, Samoa_086113.txt (+17) | Samoa_086113.txt, Gerald_Reive_c5cd23.txt | ✅ |
| q59 | Tecumseh_bcf68b.txt, Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt (+17) | Tippecanoe_order_of_battle_d27975.txt, Tecumseh_bcf68b.txt | ✅ |
| q60 | Samuel_Sim_07d71f.txt, Tromeo_and_Juliet_86fab3.txt, Bedknobs_and_Broomsticks_090f32.txt (+17) | Samuel_Sim_07d71f.txt, Awake__film__360ee6.txt | ✅ |
| q61 | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt, Øresund_Region_549ba5.txt (+17) | Øresund_Bridge_738526.txt, Øresund_Region_549ba5.txt | ✅ |
| q62 | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt, Larry_Drake_85028c.txt (+17) | Pat_Hingle_5a5cf2.txt, Clint_Eastwood_b84954.txt | ✅ |
| q63 | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt, Herbert_Akroyd_Stuart_0222bb.txt (+17) | Maurice_Ward_158a0a.txt, Starlite_2c1671.txt | ✅ |
| q64 | United_States_v._Paramount_Pictures__Inc._4aa665.txt, Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+17) | Craig_v._Boren_910f00.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt | ✅ |
| q65 | Children_s_Mercy_Park_c9283f.txt, Arrowhead_Stadium_eae21e.txt, CommunityAmerica_Ballpark_1e3f7d.txt (+17) | Children_s_Mercy_Park_c9283f.txt, CommunityAmerica_Ballpark_1e3f7d.txt | ✅ |
| q66 | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt, The_Informant__376ad6.txt (+17) | Strip_search_phone_call_scam_9b1a1b.txt, Compliance__film__3647bc.txt | ✅ |
| q67 | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt, Eucryphia_3801b3.txt (+17) | Ehretia_39d6d4.txt, Xanthoceras_306d61.txt | ✅ |
| q68 | Something_There_380707.txt, Paige_O_Hara_4492a0.txt, Something_There_380707.txt (+17) | Paige_O_Hara_4492a0.txt, Something_There_380707.txt | ✅ |
| q69 | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt, Laleli_Mosque_c7818f.txt (+17) | Nusretiye_Clock_Tower_9bcc17.txt, Eski_Imaret_Mosque_b4d62c.txt | ✅ |
| q70 | Opry_Mills_53c375.txt, Music_City_Queen_82e4a1.txt, Southaven__Mississippi_7cc376.txt (+17) | Music_City_Queen_82e4a1.txt, Opry_Mills_53c375.txt | ✅ |
| q71 | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt, Opry_Mills_53c375.txt (+17) | Spirit_Halloween_6adf3e.txt, Spencer_Gifts_c0e0e3.txt | ✅ |
| q72 | James_Fieser_e21429.txt, Berea_College_88c6f9.txt, Azusa_Pacific_University_fcee08.txt (+17) | James_Fieser_e21429.txt, Berea_College_88c6f9.txt | ✅ |
| q73 | James_Burke__science_historian__6fe4bf.txt, Connections__TV_series__deb419.txt, Connections__TV_series__deb419.txt (+17) | Connections__TV_series__deb419.txt, James_Burke__science_historian__6fe4bf.txt | ✅ |
| q74 | Romeo_87e7c1.txt, Benvolio_848ddf.txt, The_Magical_Legend_of_the_Leprechauns_0d732f.txt (+17) | Romeo_87e7c1.txt, Benvolio_848ddf.txt | ✅ |
| q75 | Addis_Ababa_6c835e.txt, National_Archives_and_Library_of_Ethiopia_6068a1.txt, Ogallala__Nebraska_3578d6.txt (+17) | National_Archives_and_Library_of_Ethiopia_6068a1.txt, Addis_Ababa_6c835e.txt | ✅ |
| q76 | Night_Ferry__composition__9f4c6a.txt, Toshi_Ichiyanagi_e50d8d.txt, Symphony_Center_f29c57.txt (+17) | Night_Ferry__composition__9f4c6a.txt, Symphony_Center_f29c57.txt | ✅ |
| q77 | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt, Laura_Osnes_4365e8.txt (+17) | Grease__The_New_Broadway_Cast_Recording_f7ac05.txt, Laura_Osnes_4365e8.txt | ✅ |
| q78 | Eucryphia_3801b3.txt, Lepidozamia_4d1d3c.txt, Elatostema_fff908.txt (+17) | Lepidozamia_4d1d3c.txt, Eucryphia_3801b3.txt | ✅ |
| q79 | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt, Samoa_086113.txt (+17) | Butch_Van_Artsdalen_be35e1.txt, Waimea_Bay__Hawaii_995124.txt | ✅ |
| q80 | Kris_Marshall_97e2d7.txt, Death_in_Paradise__TV_series__8a650d.txt, Eddie_Izzard_fb2be0.txt (+17) | Death_in_Paradise__TV_series__8a650d.txt, Kris_Marshall_97e2d7.txt | ✅ |
| q81 | EgyptAir_Flight_990_dfd74a.txt, Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt (+17) | Gameel_Al-Batouti_e3ff1b.txt, EgyptAir_Flight_990_dfd74a.txt | ✅ |
| q82 | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt, Oz_the_Great_and_Powerful_510b50.txt (+17) | Sacred_Planet_26fd7b.txt, Oz_the_Great_and_Powerful_510b50.txt | ✅ |
| q83 | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt, Jacques_Sernas_2c77cc.txt (+17) | Henry_III_of_France_2b1ba3.txt, Jean_Baptiste_Androuet_du_Cerceau_735d16.txt | ✅ |
| q84 | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt, Samoa_086113.txt (+17) | Church_of_the_Guanche_People_345bf1.txt, Tenerife_c81266.txt | ✅ |
| q85 | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt, Second_Anglo-Afghan_War_adb0b2.txt (+17) | Second_Anglo-Afghan_War_adb0b2.txt, Treaty_of_Gandamak_9d4d6d.txt | ✅ |
| q86 | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt, Bolton_6ce6c9.txt (+17) | Rivington_Hall_Barn_ad6e1c.txt, Bolton_6ce6c9.txt | ✅ |
| q87 | Hot_air_engine_2ed8e1.txt, Herbert_Akroyd_Stuart_0222bb.txt, George_Cayley_2c8397.txt (+17) | George_Cayley_2c8397.txt, Hot_air_engine_2ed8e1.txt | ✅ |
| q88 | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt, Beauty_and_the_Beast__1991_film__d38192.txt (+17) | Leonberger_d19b69.txt, Basset_Hound_ca5229.txt | ✅ |
| q89 | Northumbrian_dialect_7334ca.txt, Kingdom_of_Northumbria_4036b8.txt, Kingdom_of_the_Isles_9a036f.txt (+17) | Kingdom_of_Northumbria_4036b8.txt, Northumbrian_dialect_7334ca.txt | ✅ |
| q90 | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt, Southaven__Mississippi_7cc376.txt (+17) | Lake_Louisvilla__Louisville_799a9a.txt, Oldham_County__Kentucky_34ed53.txt | ✅ |
| q91 | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt, 1995_96_Manchester_United_F.C._season_3f9756.txt (+17) | Liu_Ailing_cfd610.txt, FIFA_Women_s_World_Cup_48fba6.txt | ✅ |
| q92 | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+17) | Rock_Springs__short_story_collection__59545d.txt, Richard_Ford_db80e7.txt | ✅ |
| q93 | Oedipus_Rex_d47dfb.txt, Dostoevsky_and_Parricide_f04c2c.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+17) | Dostoevsky_and_Parricide_f04c2c.txt, Oedipus_Rex_d47dfb.txt | ✅ |
| q94 | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt, Chelsea_Peretti_198d48.txt (+17) | Katherine_Waterston_adb28e.txt, Chrisann_Brennan_25ed97.txt | ✅ |
| q95 | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt, Kunming_08b9a1.txt (+17) | Kunming_08b9a1.txt, Yunnan_Provincial_Museum_c8f9fc.txt | ✅ |
| q96 | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt, United_States_v._Paramount_Pictures__Inc._4aa665.txt (+17) | Reynolds_v._Sims_31c83e.txt, Selle_v._Gibb_676b22.txt | ✅ |
| q97 | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt, Adnan_Akmal_e863b9.txt (+17) | Kamran_Akmal_9f2a8d.txt, Adnan_Akmal_e863b9.txt | ✅ |
| q98 | Arrowhead_Stadium_eae21e.txt, Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt (+17) | Charles_Deaton_e4fe98.txt, Arrowhead_Stadium_eae21e.txt | ✅ |
| q99 | Happy_Death_Day_793b37.txt, Jessica_Rothe_59deb4.txt, The_Bye_Bye_Man_906e6d.txt (+17) | Jessica_Rothe_59deb4.txt, Happy_Death_Day_793b37.txt | ✅ |
| q100 | Garden_Island_Naval_Chapel_6a3c4e.txt, Royal_Australian_Navy_5e1d16.txt, Samoa_086113.txt (+17) | Royal_Australian_Navy_5e1d16.txt, Garden_Island_Naval_Chapel_6a3c4e.txt | ✅ |
| q101 | The_Informant__376ad6.txt, Mark_Whitacre_ccc607.txt, Awake__film__360ee6.txt (+17) | Mark_Whitacre_ccc607.txt, The_Informant__376ad6.txt | ✅ |
| q102 | Current_Mood_0317b8.txt, Small_Town_Boy__song__d2fddb.txt, Based_on_a_True_Story..._42790b.txt (+17) | Small_Town_Boy__song__d2fddb.txt, Current_Mood_0317b8.txt | ✅ |
| q103 | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt, Impresario_5cc7af.txt (+17) | Impresario_5cc7af.txt, Vanessa_Bley_785093.txt | ✅ |
| q104 | Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt, Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt (+17) | Gargoyles__TV_series__acb424.txt, Gargoyles_the_Movie__The_Heroes_Awaken_5c24d6.txt | ✅ |
| q105 | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt, Tim_Hecker_80429a.txt (+17) | Tim_Hecker_80429a.txt, Ravedeath__1972_08847b.txt | ✅ |
| q106 | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt, Adrian_Peterson_1ca55b.txt (+17) | Ivory_Lee_Brown_4d1964.txt, Adrian_Peterson_1ca55b.txt | ✅ |
| q107 | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt, Goetheanum_5744ff.txt (+17) | Jens_Risom_59bab6.txt, Scandinavian_design_4ad552.txt | ✅ |
| q108 | The_Ganymede_Takeover_5feff5.txt, The_Man_in_the_High_Castle_1c064a.txt, Wendell_Berry_4313ab.txt (+17) | The_Man_in_the_High_Castle_1c064a.txt, The_Ganymede_Takeover_5feff5.txt | ✅ |
| q109 | Curt_Menefee_0ab3d9.txt, Michael_Strahan_1fa88f.txt, Michael_Strahan_1fa88f.txt (+17) | Michael_Strahan_1fa88f.txt, Curt_Menefee_0ab3d9.txt | ✅ |
| q110 | Summer_of_the_Monkeys_0b84ea.txt, William_Allen_White_d1418b.txt, T._R._M._Howard_0bb121.txt (+17) | William_Allen_White_d1418b.txt, Summer_of_the_Monkeys_0b84ea.txt | ✅ |
| q111 | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt, Shoba_Chandrasekhar_5285b2.txt (+17) | War_Chhod_Na_Yaar_98ceca.txt, Soha_Ali_Khan_b0280e.txt | ✅ |
| q112 | Alice_Upside_Down_2fe728.txt, Lucy_Fry_218cb3.txt, Caroline_Carver__actress__7577ed.txt (+17) | Lucas_Grabeel_c00cb8.txt, Alice_Upside_Down_2fe728.txt | ✅ |
| q113 | Snowdrop__game_engine__750d41.txt, Tom_Clancy_s_The_Division_3e09f9.txt, Icehouse_pieces_5e75b5.txt (+17) | Tom_Clancy_s_The_Division_3e09f9.txt, Snowdrop__game_engine__750d41.txt | ✅ |
| q114 | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt, Tom_Clancy_s_The_Division_3e09f9.txt (+17) | Kill_Doctor_Lucky_9b6f89.txt, Icehouse_pieces_5e75b5.txt | ✅ |
| q115 | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt, Jean-Loup_Chrétien_1d1a07.txt (+17) | Jacques_Sernas_2c77cc.txt, Fugitive_in_Trieste_041e68.txt | ✅ |
| q116 | Banshee_5e6ebd.txt, VMAQT-1_05fe33.txt, VMAQT-1_05fe33.txt (+17) | VMAQT-1_05fe33.txt, Banshee_5e6ebd.txt | ✅ |
| q117 | Barbara_Niven_fbf739.txt, Awake__film__360ee6.txt, Alice_Upside_Down_2fe728.txt (+17) | Dead_at_17_273b88.txt, Barbara_Niven_fbf739.txt | ✅ |
| q118 | Bart_the_Fink_77de9b.txt, Krusty_the_Clown_3e3656.txt, Cedric_the_Entertainer_e39d6e.txt (+17) | Krusty_the_Clown_3e3656.txt, Bart_the_Fink_77de9b.txt | ✅ |
| q119 | Viaport_Rotterdam_760a19.txt, Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt (+17) | Viaport_Rotterdam_760a19.txt, Wilton_Mall_66971e.txt | ✅ |
| q120 | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt, Marco_Da_Silva__dancer__777e91.txt (+17) | The_Muppet_Christmas_Carol_81e722.txt, Bedknobs_and_Broomsticks_090f32.txt | ✅ |
| q121 | Ambrose_Mendy_eddce7.txt, Chris_Eubank_Jr._82fe88.txt, Peter_Schmeichel_0b144a.txt (+17) | Chris_Eubank_Jr._82fe88.txt, Ambrose_Mendy_eddce7.txt | ✅ |
| q122 | Allen__amp__Company_Sun_Valley_Conference_083cbf.txt, Rupert_Murdoch_8801f1.txt, Joe_Scarborough_fde209.txt (+17) | Rupert_Murdoch_8801f1.txt, Allen__amp__Company_Sun_Valley_Conference_083cbf.txt | ✅ |
| q123 | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt, Raymond_Ochoa_da4d56.txt (+17) | Larry_Drake_85028c.txt, Gryphon__film__f811a3.txt | ✅ |
| q124 | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt, Phoenix_Television_29103f.txt (+17) | Phoenix_Television_29103f.txt, Phoenix_Hong_Kong_Channel_dbe55e.txt | ✅ |
| q125 | Patricia_Longo_b7fcef.txt, Graduados_7592c5.txt, Tenerife_c81266.txt (+17) | Graduados_7592c5.txt, Patricia_Longo_b7fcef.txt | ✅ |
| q126 | Ogallala_Aquifer_a2b49c.txt, Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt (+17) | Ogallala_Aquifer_a2b49c.txt, Ogallala__Nebraska_3578d6.txt | ✅ |
| q127 | Blinding_Edge_Pictures_b8de5a.txt, Unbreakable__film__52d8de.txt, Tron_8f60c9.txt (+17) | Unbreakable__film__52d8de.txt, Blinding_Edge_Pictures_b8de5a.txt | ✅ |
| q128 | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt, The_Good_Dinosaur_170ac4.txt (+17) | Raymond_Ochoa_da4d56.txt, The_Good_Dinosaur_170ac4.txt | ✅ |
| q129 | BraveStarr_8d412d.txt, Celebrity_Home_Entertainment_c01bf6.txt, Gargoyles__TV_series__acb424.txt (+17) | Celebrity_Home_Entertainment_c01bf6.txt, BraveStarr_8d412d.txt | ✅ |
| q130 | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt, The_Informant__376ad6.txt (+17) | Kam_Heskin_c848f5.txt, The_Prince_and_Me_253bec.txt | ✅ |
| q131 | Lucy_Fry_218cb3.txt, Jessica_Rothe_59deb4.txt, Nina_Dobrev_05e14f.txt (+17) | Mr._Church_ce0d51.txt, Lucy_Fry_218cb3.txt | ✅ |
| q132 | Shoba_Chandrasekhar_5285b2.txt, Ithu_Engal_Neethi_ad89e5.txt, Soha_Ali_Khan_b0280e.txt (+17) | Ithu_Engal_Neethi_ad89e5.txt, Shoba_Chandrasekhar_5285b2.txt | ✅ |
| q133 | Official_Ireland_2ed543.txt, Catholic_Church_in_Ireland_77cac6.txt, Catholic_Church_in_Ireland_77cac6.txt (+17) | Catholic_Church_in_Ireland_77cac6.txt, Official_Ireland_2ed543.txt | ✅ |
| q134 | Bridge_to_Terabithia__1985_film__1aaa6c.txt, Bedknobs_and_Broomsticks_090f32.txt, Bridge_to_Terabithia__novel__21de92.txt (+17) | Bridge_to_Terabithia__novel__21de92.txt, Bridge_to_Terabithia__1985_film__1aaa6c.txt | ✅ |
| q135 | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt, Curt_Menefee_0ab3d9.txt (+17) | Joe_Scarborough_fde209.txt, Morning_Joe_a4a08d.txt | ✅ |
| q136 | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt, Kasper_Schmeichel_c9da28.txt (+17) | Scout_Tufankjian_755d73.txt, Art_Laboe_572809.txt | ✅ |
| q137 | Atari_Assembler_Editor_a5f9fc.txt, Shepardson_Microsystems_0fa820.txt, Snowdrop__game_engine__750d41.txt (+17) | Shepardson_Microsystems_0fa820.txt, Atari_Assembler_Editor_a5f9fc.txt | ✅ |
| q138 | His_Band_and_the_Street_Choir_f5b88a.txt, I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt (+17) | I_ve_Been_Working_b0df7c.txt, His_Band_and_the_Street_Choir_f5b88a.txt | ✅ |
| q139 | Aldosterone_b32476.txt, Aldosterone_b32476.txt, Angiotensin_3f2772.txt (+17) | Angiotensin_3f2772.txt, Aldosterone_b32476.txt | ✅ |
| q140 | Nancy_Soderberg_a37451.txt, United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt (+17) | United_States_elections__2018_95c5df.txt, Nancy_Soderberg_a37451.txt | ✅ |
| q141 | Mississippi_University_for_Women_v._Hogan_f450c0.txt, Berghuis_v._Thompkins_c16323.txt, Reynolds_v._Sims_31c83e.txt (+17) | Berghuis_v._Thompkins_c16323.txt, Mississippi_University_for_Women_v._Hogan_f450c0.txt | ✅ |
| q142 | AFC_North_8f76d9.txt, AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt (+17) | AFC_North_8f76d9.txt, 2009_Cleveland_Browns_season_f9d029.txt | ✅ |
| q143 | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt (+17) | Dorothea_Jordan_1695d3.txt, Elizabeth_Hay__Countess_of_Erroll_6187ae.txt | ✅ |
| q144 | Honda_Ballade_4d8914.txt, Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt (+17) | Honda_CR-X_9602f7.txt, Honda_Ballade_4d8914.txt | ✅ |
| q145 | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt, The_Monster__song__886456.txt (+17) | Numb__Rihanna_song__6c9d1e.txt, The_Monster__song__886456.txt | ✅ |
| q146 | Backstage__magazine__258d77.txt, Celebrity_Home_Entertainment_c01bf6.txt, Christopher_Oscar_Peña_3991f4.txt (+17) | Backstage__magazine__258d77.txt, Christopher_Oscar_Peña_3991f4.txt | ✅ |
| q147 | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt, Music_for_Electric_Metronomes_0c93d9.txt (+17) | Toshi_Ichiyanagi_e50d8d.txt, Music_for_Electric_Metronomes_0c93d9.txt | ✅ |
| q148 | Ego_the_Living_Planet_f6f847.txt, Guardians_of_the_Galaxy_Vol._2_b6c488.txt, James_Gunn_9a06f8.txt (+17) | Guardians_of_the_Galaxy_Vol._2_b6c488.txt, Ego_the_Living_Planet_f6f847.txt | ✅ |
| q149 | Sponsorship_scandal_5e9ed3.txt, Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt (+17) | Government_of_Canada_46a20a.txt, Sponsorship_scandal_5e9ed3.txt | ✅ |
| q150 | Conrad_Anker_9d25ca.txt, George_Mallory_ccb70e.txt, Apsley_Cherry-Garrard_ff4e03.txt (+17) | George_Mallory_ccb70e.txt, Conrad_Anker_9d25ca.txt | ✅ |
| q151 | Joan_Crawford_5fcb43.txt, The_Duke_Steps_Out_7cc9f2.txt, Erika_Jayne_60e347.txt (+17) | The_Duke_Steps_Out_7cc9f2.txt, Joan_Crawford_5fcb43.txt | ✅ |
| q152 | Carol__film__cbfb1f.txt, Guild_of_Music_Supervisors_Awards_7984fd.txt, The_Muppet_Christmas_Carol_81e722.txt (+17) | Guild_of_Music_Supervisors_Awards_7984fd.txt, Carol__film__cbfb1f.txt | ✅ |
| q153 | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+17) | Richard_Hornsby__amp__Sons_fb744c.txt, Herbert_Akroyd_Stuart_0222bb.txt | ✅ |
| q154 | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt, Bishop_Carroll_Catholic_High_School_5aa312.txt (+17) | Cardinal_Pole_Roman_Catholic_School_3d0515.txt, Reginald_Pole_a30667.txt | ✅ |
| q155 | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt, Goetheanum_5744ff.txt (+17) | Jörgen_Smit_869bf0.txt, Goetheanum_5744ff.txt | ✅ |
| q156 | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt, Berea_College_88c6f9.txt (+17) | Associate_degree_f57219.txt, Southeastern_Illinois_College_22bfe4.txt | ✅ |
| q157 | Sprawl_trilogy_894d39.txt, Neuromancer_6029bf.txt, The_Man_in_the_High_Castle_1c064a.txt (+17) | Neuromancer_6029bf.txt, Sprawl_trilogy_894d39.txt | ✅ |
| q158 | Azad_Hind_Dal_a97cda.txt, Subhas_Chandra_Bose_4b75c6.txt, Second_Anglo-Afghan_War_adb0b2.txt (+17) | Subhas_Chandra_Bose_4b75c6.txt, Azad_Hind_Dal_a97cda.txt | ✅ |
| q159 | Eddie_Izzard_fb2be0.txt, Stripped__tour__91c664.txt, Chelsea_Peretti_198d48.txt (+17) | Stripped__tour__91c664.txt, Eddie_Izzard_fb2be0.txt | ✅ |
| q160 | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt, Soha_Ali_Khan_b0280e.txt (+17) | Hiren_Roy_9d83f5.txt, Vilayat_Khan_2d1c60.txt | ✅ |
| q161 | Ecballium_ad2282.txt, Elatostema_fff908.txt, Thalictrum_3222b5.txt (+17) | Elatostema_fff908.txt, Ecballium_ad2282.txt | ✅ |
| q162 | Polypodium_71abb6.txt, Aichryson_319a15.txt, Thalictrum_3222b5.txt (+17) | Polypodium_71abb6.txt, Aichryson_319a15.txt | ✅ |
| q163 | Adoption_2002_46d793.txt, Adoption_and_Safe_Families_Act_b4c27b.txt, Conscription_in_the_United_States_e920fa.txt (+17) | Adoption_and_Safe_Families_Act_b4c27b.txt, Adoption_2002_46d793.txt | ✅ |
| q164 | Crash_Pad_ed95da.txt, Nina_Dobrev_05e14f.txt, The_Lodge__TV_series__b905ce.txt (+17) | Nina_Dobrev_05e14f.txt, Crash_Pad_ed95da.txt | ✅ |
| q165 | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt, Their_Lives_eed122.txt (+17) | The_Company_They_Keep_a91182.txt, Azusa_Pacific_University_fcee08.txt | ✅ |
| q166 | Parks_and_Recreation_051519.txt, Ms._Knope_Goes_to_Washington_799ad4.txt, Will__amp__Grace_31b102.txt (+17) | Ms._Knope_Goes_to_Washington_799ad4.txt, Parks_and_Recreation_051519.txt | ✅ |
| q167 | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt, Kasper_Schmeichel_c9da28.txt (+17) | 2011_La_Manga_Cup_3d15dc.txt, IK_Start_d60e98.txt | ✅ |
| q168 | Can_t_Fight_the_Moonlight_caf06b.txt, The_Best_of_LeAnn_Rimes_722f10.txt, Tron_8f60c9.txt (+17) | The_Best_of_LeAnn_Rimes_722f10.txt, Can_t_Fight_the_Moonlight_caf06b.txt | ✅ |
| q169 | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt, Lee_Seung-gi_1ed0e1.txt (+17) | My_Secret_Hotel_6b2b41.txt, Yoo_In-na_5b538a.txt | ✅ |
| q170 | Unbelievable__The_Notorious_B.I.G._song__91da75.txt, Ready_to_Die_c706b2.txt, Ready_to_Die_c706b2.txt (+17) | Ready_to_Die_c706b2.txt, Unbelievable__The_Notorious_B.I.G._song__91da75.txt | ✅ |
| q171 | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt, Cedric_the_Entertainer_e39d6e.txt (+17) | Fantasy_Records_8193ed.txt, Vince_Guaraldi_20d4e1.txt | ✅ |
| q172 | Violin_Sonata_No._5__Beethoven__8f6942.txt, Symphony_No._7__Beethoven__9c3b01.txt, Hugo_von_Hofmannsthal_cd4c9e.txt (+17) | Symphony_No._7__Beethoven__9c3b01.txt, Violin_Sonata_No._5__Beethoven__8f6942.txt | ✅ |
| q173 | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt, Mondelez_International_9a57f2.txt (+17) | House_of_Pies_9fe39f.txt, Orange_Julius_e106c8.txt | ✅ |
| q174 | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt, Cedric_the_Entertainer_e39d6e.txt (+17) | Black_Movie_Awards_9161c5.txt, Cedric_the_Entertainer_e39d6e.txt | ✅ |
| q175 | Saab_36_108e14.txt, Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt (+17) | Saab_36_108e14.txt, Avro_Vulcan_7aa981.txt | ✅ |
| q176 | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt, Pasek_and_Paul_aa5312.txt (+17) | Arizona__song__02210d.txt, Mark_Lindsay_c8bd25.txt | ✅ |
| q177 | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt, North_American_T-6_Texan_76936b.txt (+17) | North_American_Aviation_377efe.txt, North_American_T-6_Texan_76936b.txt | ✅ |
| q178 | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt, T._R._M._Howard_0bb121.txt (+17) | David_T._Beito_7da5a5.txt, T._R._M._Howard_0bb121.txt | ✅ |
| q179 | Lucas_Carvalho_940f12.txt, 4___400_metres_relay_8bba84.txt, 4___400_metres_relay_8bba84.txt (+17) | 4___400_metres_relay_8bba84.txt, Lucas_Carvalho_940f12.txt | ✅ |
| q180 | Watercliffe_Meadow_Community_Primary_School_1e4bca.txt, Political_correctness_f6d0c6.txt, Cardinal_Pole_Roman_Catholic_School_3d0515.txt (+17) | Political_correctness_f6d0c6.txt, Watercliffe_Meadow_Community_Primary_School_1e4bca.txt | ✅ |
| q181 | Parodia_a9ceb8.txt, Thalictrum_3222b5.txt, Polypodium_71abb6.txt (+17) | Thalictrum_3222b5.txt, Parodia_a9ceb8.txt | ✅ |
| q182 | The_Simpsons__season_23__36cc7b.txt, At_Long_Last_Leave_8151f2.txt, Bart_the_Fink_77de9b.txt (+17) | At_Long_Last_Leave_8151f2.txt, The_Simpsons__season_23__36cc7b.txt | ✅ |
| q183 | BJ_s_Wholesale_Club_f114e3.txt, US_Vision_f98d42.txt, US_Vision_f98d42.txt (+17) | US_Vision_f98d42.txt, BJ_s_Wholesale_Club_f114e3.txt | ✅ |
| q184 | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt, Lochmere_Archeological_District_2ae282.txt (+17) | Pennacook_2fae97.txt, Lochmere_Archeological_District_2ae282.txt | ✅ |
| q185 | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt, Southaven__Mississippi_7cc376.txt (+17) | Peabody_Hotel_73c038.txt, Hyatt_Regency_Orlando_e28755.txt | ✅ |
| q186 | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt, Samantha_Cristoforetti_fc1bcd.txt (+17) | Samantha_Cristoforetti_fc1bcd.txt, Jean-Loup_Chrétien_1d1a07.txt | ✅ |
| q187 | Flynn_Rider_37adf2.txt, Zachary_Levi_9c18c7.txt, Lucas_Grabeel_c00cb8.txt (+17) | Zachary_Levi_9c18c7.txt, Flynn_Rider_37adf2.txt | ✅ |
| q188 | Vacation_with_Derek_da8ab6.txt, Life_with_Derek_547274.txt, Cressida_Bonas_5c3c57.txt (+17) | Life_with_Derek_547274.txt, Vacation_with_Derek_da8ab6.txt | ✅ |
| q189 | Jin_Jing_fa2e08.txt, Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt (+17) | Carrefour_a2a75b.txt, Jin_Jing_fa2e08.txt | ✅ |
| q190 | The_Worst_Journey_in_the_World_1fd01f.txt, Apsley_Cherry-Garrard_ff4e03.txt, Apsley_Cherry-Garrard_ff4e03.txt (+17) | Apsley_Cherry-Garrard_ff4e03.txt, The_Worst_Journey_in_the_World_1fd01f.txt | ✅ |
| q191 | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt, James_Fieser_e21429.txt (+17) | Vincent_Kling__translator__0babbb.txt, Hugo_von_Hofmannsthal_cd4c9e.txt | ✅ |
| q192 | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, The_Prince_and_Me_253bec.txt, Lee_Seung-gi_1ed0e1.txt (+17) | My_Girlfriend_Is_a_Nine-Tailed_Fox_840a28.txt, Lee_Seung-gi_1ed0e1.txt | ✅ |
| q193 | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt, William_Lever__1st_Viscount_Leverhulme_e60438.txt (+17) | William_Lever__1st_Viscount_Leverhulme_e60438.txt, Lady_Lever_Art_Gallery_fc8842.txt | ✅ |
| q194 | Thomas_Doherty__actor__7caf88.txt, The_Lodge__TV_series__b905ce.txt, Life_with_Derek_547274.txt (+17) | The_Lodge__TV_series__b905ce.txt, Thomas_Doherty__actor__7caf88.txt | ✅ |
| q195 | Chelsea_Peretti_198d48.txt, Brooklyn_Nine-Nine_eec7c9.txt, Larry_Drake_85028c.txt (+17) | Brooklyn_Nine-Nine_eec7c9.txt, Chelsea_Peretti_198d48.txt | ✅ |
| q196 | Mark_Gaudet_0f4aa3.txt, Jan_Axel_Blomberg_a42602.txt, Mark_Lindsay_c8bd25.txt (+17) | Jan_Axel_Blomberg_a42602.txt, Mark_Gaudet_0f4aa3.txt | ✅ |
| q197 | Miami_Canal_161997.txt, Dundee_Canal_05ea0b.txt, Richard_Hornsby__amp__Sons_fb744c.txt (+17) | Dundee_Canal_05ea0b.txt, Miami_Canal_161997.txt | ✅ |
| q198 | Tron_8f60c9.txt, The_Million_Dollar_Duck_d1d45c.txt, Gryphon__film__f811a3.txt (+17) | The_Million_Dollar_Duck_d1d45c.txt, Tron_8f60c9.txt | ✅ |
| q199 | Beauty_and_the_Beast__franchise__7a780a.txt, Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt (+17) | Beauty_and_the_Beast__1991_film__d38192.txt, Beauty_and_the_Beast__franchise__7a780a.txt | ✅ |
| q200 | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt, Peter_Schmeichel_0b144a.txt (+17) | Dave_Schultz__wrestling__7d8c52.txt, Foxcatcher_874805.txt | ✅ |

---

## Metrics Guide

- **Recall@K**: Fraction of relevant docs found in top-K results (higher is better)
- **MRR**: Mean Reciprocal Rank — how high the first relevant result ranks (higher is better)
- **Coverage**: Fraction of all relevant docs ever retrieved across queries (higher is better)
- **Redundancy**: Average times each doc is retrieved (lower may indicate diverse results)

---

*Generated by [RagTune](https://github.com/metawake/ragtune)*
