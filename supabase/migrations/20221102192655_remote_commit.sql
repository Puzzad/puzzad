-- This script was generated by the Schema Diff utility in pgAdmin 4
-- For the circular dependencies, the order in which Schema Diff writes the objects is not very sophisticated
-- and may require manual changes to the script to ensure changes are applied in the correct order.
-- Please report an issue for any failure with the reproduction steps.

-- Extension: pgaudit

-- DROP EXTENSION pgaudit;

CREATE EXTENSION IF NOT EXISTS pgaudit
    SCHEMA extensions
    VERSION "1.6.1";

CREATE OR REPLACE FUNCTION public.finishadventure()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
  if NEW.puzzle is null then
    NEW.status = 'EXPIRED';
  end if;
  return NEW;
END;
$BODY$;

ALTER FUNCTION public.finishadventure()
    OWNER TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO anon;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO postgres;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO authenticated;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO service_role;

GRANT EXECUTE ON FUNCTION public.finishadventure() TO PUBLIC;

CREATE OR REPLACE FUNCTION public.checkguess()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF SECURITY DEFINER
AS $BODY$
BEGIN
  NEW.correct = (select count(*) > 0 FROM answers WHERE puzzle = NEW.puzzle AND answer = NEW.content);
  RETURN NEW;
END;
$BODY$;

ALTER FUNCTION public.checkguess()
    OWNER TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.checkguess() TO anon;

GRANT EXECUTE ON FUNCTION public.checkguess() TO postgres;

GRANT EXECUTE ON FUNCTION public.checkguess() TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.checkguess() TO authenticated;

GRANT EXECUTE ON FUNCTION public.checkguess() TO service_role;

GRANT EXECUTE ON FUNCTION public.checkguess() TO PUBLIC;

CREATE OR REPLACE FUNCTION public.advancetonextpuzzle()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF SECURITY DEFINER
AS $BODY$
BEGIN
  if NEW.correct then
    UPDATE games SET puzzle = (SELECT next FROM puzzles WHERE id = NEW.puzzle) WHERE games.code = NEW.game;
    return NEW;
  end if;
RETURN NULL;    
END;
$BODY$;

ALTER FUNCTION public.advancetonextpuzzle()
    OWNER TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO anon;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO postgres;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO supabase_admin;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO authenticated;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO service_role;

GRANT EXECUTE ON FUNCTION public.advancetonextpuzzle() TO PUBLIC;

CREATE OR REPLACE FUNCTION public.setteamcode()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    VOLATILE
    COST 100
AS $BODY$
BEGIN
  NEW.code = (select CONCAT((select value as adjective from adjectives order by random() limit 1), '-', (select value as colour from colours order by random() limit 1), '-', (select value as animal from animals order by random() limit 1)) as code);
  RETURN NEW;
END;
$BODY$;

CREATE TABLE IF NOT EXISTS public.answers
(
    id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    puzzle bigint,
    answer character varying COLLATE pg_catalog."default",
    CONSTRAINT answers_pkey PRIMARY KEY (id),
    CONSTRAINT answers_puzzle_fkey FOREIGN KEY (puzzle)
        REFERENCES public.puzzles (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.answers
    OWNER to supabase_admin;

ALTER TABLE IF EXISTS public.answers
    ENABLE ROW LEVEL SECURITY;

GRANT ALL ON TABLE public.answers TO anon;

GRANT ALL ON TABLE public.answers TO postgres;

GRANT ALL ON TABLE public.answers TO supabase_admin;

GRANT ALL ON TABLE public.answers TO authenticated;

GRANT ALL ON TABLE public.answers TO service_role;

ALTER TABLE IF EXISTS public.guesses DROP COLUMN IF EXISTS "user";

ALTER TABLE IF EXISTS public.guesses
    ADD COLUMN game character varying COLLATE pg_catalog."default" NOT NULL;

ALTER TABLE IF EXISTS public.guesses
    ADD COLUMN correct boolean DEFAULT false;
ALTER TABLE IF EXISTS public.guesses DROP CONSTRAINT IF EXISTS guesses_user_fkey;

ALTER TABLE IF EXISTS public.guesses
    ADD CONSTRAINT guesses_game_fkey FOREIGN KEY (game)
    REFERENCES public.games (code) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

CREATE POLICY "Read guesses for owned games"
    ON public.guesses
    AS PERMISSIVE
    FOR SELECT
    TO public
    USING ((auth.uid() = ( SELECT games."user"
   FROM games
  WHERE ((games.code)::text = (guesses.game)::text))));

CREATE POLICY "Write guess for current puzzle on owned game"
    ON public.guesses
    AS PERMISSIVE
    FOR INSERT
    TO public
    WITH CHECK ((auth.uid() = ( SELECT games."user"
   FROM games
  WHERE (((games.code)::text = (guesses.game)::text) AND (games.puzzle = guesses.puzzle)))));

CREATE TRIGGER "AdvanceToNextPuzzle"
    AFTER INSERT
    ON public.guesses
    FOR EACH ROW
    EXECUTE FUNCTION public.advancetonextpuzzle();
CREATE TRIGGER "CheckGuess"
    BEFORE INSERT
    ON public.guesses
    FOR EACH ROW
    EXECUTE FUNCTION public.checkguess();

REVOKE ALL ON TABLE public.adjectives FROM anon;
REVOKE ALL ON TABLE public.adjectives FROM authenticated;
REVOKE ALL ON TABLE public.adjectives FROM service_role;
REVOKE ALL ON TABLE public.adjectives FROM supabase_admin;
GRANT ALL ON TABLE public.adjectives TO anon;

GRANT ALL ON TABLE public.adjectives TO supabase_admin;

GRANT ALL ON TABLE public.adjectives TO authenticated;

GRANT ALL ON TABLE public.adjectives TO service_role;

REVOKE ALL ON TABLE public.colours FROM anon;
REVOKE ALL ON TABLE public.colours FROM authenticated;
REVOKE ALL ON TABLE public.colours FROM service_role;
REVOKE ALL ON TABLE public.colours FROM supabase_admin;
GRANT ALL ON TABLE public.colours TO anon;

GRANT ALL ON TABLE public.colours TO supabase_admin;

GRANT ALL ON TABLE public.colours TO authenticated;

GRANT ALL ON TABLE public.colours TO service_role;

REVOKE ALL ON TABLE public.puzzles FROM anon;
REVOKE ALL ON TABLE public.puzzles FROM authenticated;
REVOKE ALL ON TABLE public.puzzles FROM postgres;
REVOKE ALL ON TABLE public.puzzles FROM service_role;
REVOKE ALL ON TABLE public.puzzles FROM supabase_admin;
GRANT ALL ON TABLE public.puzzles TO anon;

GRANT ALL ON TABLE public.puzzles TO postgres;

GRANT ALL ON TABLE public.puzzles TO supabase_admin;

GRANT ALL ON TABLE public.puzzles TO authenticated;

GRANT ALL ON TABLE public.puzzles TO service_role;

ALTER TABLE IF EXISTS public.puzzles DROP COLUMN IF EXISTS answer;

ALTER TABLE IF EXISTS public.puzzles DROP COLUMN IF EXISTS "order";

ALTER TABLE IF EXISTS public.puzzles
    RENAME id TO content;

ALTER TABLE public.puzzles
    ALTER COLUMN content TYPE text COLLATE pg_catalog."default";

ALTER TABLE IF EXISTS public.puzzles
    ALTER COLUMN content DROP IDENTITY;

ALTER TABLE IF EXISTS public.puzzles
    ALTER COLUMN content SET STORAGE EXTENDED;

ALTER TABLE IF EXISTS public.puzzles
    ADD COLUMN next bigint;

ALTER TABLE IF EXISTS public.puzzles
    ADD COLUMN storage_slug character varying COLLATE pg_catalog."default" NOT NULL;

ALTER TABLE IF EXISTS public.puzzles
    ADD CONSTRAINT puzzles_next_fkey FOREIGN KEY (next)
    REFERENCES public.puzzles (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

CREATE POLICY "Read access for current puzzle for owned games"
    ON public.puzzles
    AS PERMISSIVE
    FOR SELECT
    TO public
    USING (( SELECT (count(*) > 0)
   FROM games
  WHERE ((games.adventure = puzzles.adventure) AND (games."user" = auth.uid()) AND (games.puzzle = puzzles.id))));

REVOKE ALL ON TABLE public.adventures FROM anon;
REVOKE ALL ON TABLE public.adventures FROM authenticated;
REVOKE ALL ON TABLE public.adventures FROM postgres;
REVOKE ALL ON TABLE public.adventures FROM service_role;
REVOKE ALL ON TABLE public.adventures FROM supabase_admin;
GRANT ALL ON TABLE public.adventures TO anon;

GRANT ALL ON TABLE public.adventures TO postgres;

GRANT ALL ON TABLE public.adventures TO supabase_admin;

GRANT ALL ON TABLE public.adventures TO authenticated;

GRANT ALL ON TABLE public.adventures TO service_role;

ALTER TABLE IF EXISTS public.adventures DROP COLUMN IF EXISTS preview;

ALTER TABLE IF EXISTS public.adventures DROP COLUMN IF EXISTS intro;

ALTER TABLE IF EXISTS public.adventures
    ADD COLUMN "promoBackground" text COLLATE pg_catalog."default";

ALTER TABLE IF EXISTS public.adventures
    ADD COLUMN "promoLogo" text COLLATE pg_catalog."default";

ALTER TABLE IF EXISTS public.games
    RENAME id TO puzzle;

ALTER TABLE IF EXISTS public.games
    ALTER COLUMN puzzle DROP NOT NULL;

ALTER TABLE IF EXISTS public.games
    ALTER COLUMN puzzle DROP IDENTITY;
CREATE TRIGGER "FinishAdventure"
    BEFORE UPDATE 
    ON public.games
    FOR EACH ROW
    EXECUTE FUNCTION public.finishadventure();
CREATE TRIGGER "SetTeamCode"
    BEFORE INSERT
    ON public.games
    FOR EACH ROW
    EXECUTE FUNCTION public.setteamcode();

REVOKE ALL ON TABLE public.animals FROM anon;
REVOKE ALL ON TABLE public.animals FROM authenticated;
REVOKE ALL ON TABLE public.animals FROM service_role;
REVOKE ALL ON TABLE public.animals FROM supabase_admin;
GRANT ALL ON TABLE public.animals TO anon;

GRANT ALL ON TABLE public.animals TO supabase_admin;

GRANT ALL ON TABLE public.animals TO authenticated;

GRANT ALL ON TABLE public.animals TO service_role;
