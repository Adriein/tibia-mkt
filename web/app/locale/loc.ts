import * as EN from "./en";
import * as ES from "./es";
import * as BR from "./br";
import * as PL from "./pl";

export const English = "us-US";
export const Spanish = "es-ES";
export const Portuguese = "br-BR";
export const Polish = "pl-PL";

export interface PageTranslations {
    Home: HomeTranslations;
}

export interface HomeTranslations {
    welcome: string;
}

// Define a generic type that returns the correct translation interface based on the page parameter
type TranslationForPage<P extends keyof PageTranslations> = PageTranslations[P];

export const loc: <P extends keyof PageTranslations>(lang: string, page: P) => TranslationForPage<P> =
    <P extends keyof PageTranslations>(lang: string, page: P): TranslationForPage<P> => {
        switch (lang) {
            case Spanish:
                return ES.default[page] as TranslationForPage<P>;
            case Portuguese:
                return BR.default[page] as TranslationForPage<P>;
            case Polish:
                return PL.default[page] as TranslationForPage<P>;
            default:
                return EN.default[page] as TranslationForPage<P>;
        }
    }