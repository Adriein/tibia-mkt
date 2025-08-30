import * as EN from "./en";
import * as ES from "./es";
import * as BR from "./br";
import * as PL from "./pl";

export enum Locale {
    English = "en-US",
    Spanish = "es-ES",
    Portuguese = "br-BR",
    Polish = "pl-PL",
}

export enum BeautyLocale {
    English = "en",
    Spanish = "es",
    Portuguese = "br",
    Polish = "pl",
}

export type PageTranslations = typeof EN.default;
export type HomeTranslations = typeof EN.default.Home;
export type DetailTranslations = typeof EN.default.Detail;

// Define a generic type that returns the correct translation interface based on the page parameter
type TranslationForPage<P extends keyof PageTranslations> = PageTranslations[P];

export const loc: <P extends keyof PageTranslations>(lang: string, page: P) => TranslationForPage<P> =
    <P extends keyof PageTranslations>(lang: string, page: P): TranslationForPage<P> => {
        switch (lang) {
            case Locale.Spanish:
                return ES.default[page] as TranslationForPage<P>;
            case Locale.Portuguese:
                return BR.default[page] as TranslationForPage<P>;
            case Locale.Polish:
                return PL.default[page] as TranslationForPage<P>;
            default:
                return EN.default[page] as TranslationForPage<P>;
        }
    };

export function languageConverter(lang: string): Locale {
    switch (lang) {
        case BeautyLocale.Spanish:
            return Locale.Spanish;
        case BeautyLocale.English:
            return Locale.English;
        case BeautyLocale.Portuguese:
            return Locale.Portuguese;
        case BeautyLocale.Polish:
            return Locale.Polish;
        default:
            return Locale.English
    }
}