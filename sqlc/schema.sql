CREATE TABLE "public"."eav_attribute_options" (
    "id" bigserial NOT NULL,
    "eav_attribute_code" character varying(255) NOT NULL,
    "value" character varying(1024) NOT NULL,
    "ordering" smallint NOT NULL CONSTRAINT eav_attribute_options_ordering_check CHECK (ordering >= 0),
    "is_visible" boolean NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE INDEX eav_attribute_options_idx_eav_attribute_code ON public.eav_attribute_options USING btree (eav_attribute_code);

ALTER TABLE ONLY "public"."eav_attribute_options" ADD CONSTRAINT "eav_attribute_options_fk_eav_attribute_code" FOREIGN KEY ("eav_attribute_code") REFERENCES "public"."eav_attributes" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."eav_attribute_values" (
    "user_id" bigint NOT NULL,
    "eav_attribute_code" character varying(255) NOT NULL,
    "value" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id", "eav_attribute_code")
);

CREATE INDEX eav_attribute_values_idx_user_id ON public.eav_attribute_values USING btree (user_id);

CREATE INDEX eav_attribute_values_idx_eav_attribute_code ON public.eav_attribute_values USING btree (eav_attribute_code);

ALTER TABLE ONLY "public"."eav_attribute_values" ADD CONSTRAINT "eav_attribute_code" FOREIGN KEY ("eav_attribute_code") REFERENCES "public"."eav_attributes" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE ONLY "public"."eav_attribute_values" ADD CONSTRAINT "eav_attribute_values_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."eav_attributes" (
    "code" character varying(255) NOT NULL,
    "name" character varying(255) NOT NULL,
    "value_type" smallint NOT NULL CONSTRAINT eav_attributes_value_type_check CHECK (value_type >= 0),
    "description" character varying(255) NOT NULL,
    "field_format" character varying(255),
    "regexp" character varying(255),
    "min_length" smallint CONSTRAINT eav_attributes_min_length_check CHECK (min_length >= 0),
    "max_length" smallint CONSTRAINT eav_attributes_max_length_check CHECK (max_length >= 0),
    "is_selection" boolean NOT NULL DEFAULT false,
    "is_required" boolean NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("code")
);

CREATE TABLE "public"."users" (
    "id" bigserial NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
