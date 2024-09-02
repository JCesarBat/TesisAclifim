CREATE TABLE "datos_sociales"(
    id bigserial primary key,
    "id_asociado" bigint not null unique,
    "ocupacion" varchar check ( "ocupacion" in('Jubilado','sin_empleo','trabajando') ),
    "estado_civil" varchar check("estado_civil" in('soltero','casado','divorsiado','viudo')),
    "integracion_revolucionaria" varchar check ( "integracion_revolucionaria" in('PCC','UJC','CDR','FMC','MTT','BPO','CTC','FAR','MININT','otra','ACRO') )
    );
ALTER TABLE "datos_sociales" ADD FOREIGN KEY  ("id_asociado") REFERENCES "asociado" ("id");