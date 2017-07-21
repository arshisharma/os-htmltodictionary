package main

type ScrapeElement struct {
    path []string
    index int
}

var AdvancedDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".small-brand-logo"}, 0},
    "V2" : {[]string{".small-brand-logo"}, 1},
    "V3" : {[]string{".small-brand-logo"}, 2},
    "V4" : {[]string{".small-brand-logo"}, 3},
    "V5" : {[]string{".small-brand-logo"}, 4},
    "V6" : {[]string{".small-brand-logo"}, 5},
    "V7" : {[]string{".small-brand-logo"}, 6},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "V8" : {[]string{".official-shop-promo__banner-half"}, 0},
    "V9" : {[]string{".span6", ".row-fluid", ".mt-20"}, 0},
    "V10" : {[]string{".span6", ".row-fluid", ".span6.official-shop-promo__type2 "}, 1},
    "T": {[]string{".official-promo-top-video"}, 0},
}

var AdvancedMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "V1" : {[]string{".official-slide-category", ".swiper-slide"}, 0},
    "V2" : {[]string{".official-slide-category", ".swiper-slide"}, 1},
    "V3" : {[]string{".official-slide-category", ".swiper-slide"}, 2},
    "V4" : {[]string{".official-slide-category", ".swiper-slide"}, 3},
    "V5" : {[]string{".official-slide-category", ".swiper-slide"}, 4},
    "V6" : {[]string{".official-slide-category", ".swiper-slide"}, 5},
    "V7" : {[]string{".official-slide-category", ".swiper-slide"}, 6},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "V8" : {[]string{".span12"}, 0},
    "V9" : {[]string{".span6"}, 0},
    "V10" : {[]string{".span6"}, 1},
}

var AdvancedApps = AdvancedMobile

var FashionDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".official-shop-promo__type1"}, 0},
    "V2" : {[]string{".official-shop-promo__type1"}, 1},
    "V3" : {[]string{".official-shop-promo__type1"}, 2},
    "V4" : {[]string{".official-shop-promo__type1"}, 3},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "VV4" : {[]string{".swiper-slide"}, 3},
    "V5" : {[]string{".official-shop-promo__type2"}, 0},
    "V6" : {[]string{".official-shop-promo__type2"}, 1},
    "V7" : {[]string{".official-shop-promo__type2"}, 2},
    "V8" : {[]string{".official-shop-promo__type2"}, 3},
    "B1" : {[]string{".official-shop-branding__video"}, 0},
    "B2" : {[]string{".official-shop-branding__video"}, 1},
    "B3" : {[]string{".official-shop-branding__video"}, 2},
}

var FashionMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "V1" : {[]string{".swiper-slide", ".official-shop-promo__type1"}, 0},
    "V2" : {[]string{".swiper-slide", ".official-shop-promo__type1"}, 1},
    "V3" : {[]string{".swiper-slide", ".official-shop-promo__type1"}, 2},
    "V4" : {[]string{".swiper-slide", ".official-shop-promo__type1"}, 3},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "VV4" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 3},
    "V5" : {[]string{".section-official-shop-promo__type2", ".span6"}, 0},
    "V6" : {[]string{".section-official-shop-promo__type2", ".span6"}, 1},
    "V7" : {[]string{".section-official-shop-promo__type2", ".span6"}, 2},
    "V8" : {[]string{".section-official-shop-promo__type2", ".span6"}, 3},
}

var FashionApps = FashionMobile

var SophisticatedDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "VVV1" : {[]string{".small-brand-logo"}, 0},
    "VVV2" : {[]string{".small-brand-logo"}, 1},
    "VVV3" : {[]string{".small-brand-logo"}, 2},
    "VVV4" : {[]string{".small-brand-logo"}, 3},
    "VVV5" : {[]string{".small-brand-logo"}, 4},
    "VVV6" : {[]string{".small-brand-logo"}, 5},
    "VVV7" : {[]string{".small-brand-logo"}, 6},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "VV4" : {[]string{".swiper-slide"}, 3},
    "V1" : {[]string{".official-shop-promo__type2"}, 0},
    "V2" : {[]string{".official-shop-promo__type2"}, 1},
    "V3" : {[]string{".official-shop-promo__type2"}, 2},
    "V4" : {[]string{".span6.official-shop-promo__banner-half"}, 0},
    "V5" : {[]string{".span6.official-shop-promo__banner-half"}, 1},
}

var SophisticatedMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "VVV1" : {[]string{".official-shop-brands", ".swiper-slide"}, 0},
    "VVV2" : {[]string{".official-shop-brands", ".swiper-slide"}, 1},
    "VVV3" : {[]string{".official-shop-brands", ".swiper-slide"}, 2},
    "VVV4" : {[]string{".official-shop-brands", ".swiper-slide"}, 3},
    "VVV5" : {[]string{".official-shop-brands", ".swiper-slide"}, 4},
    "VVV6" : {[]string{".official-shop-brands", ".swiper-slide"}, 5},
    "VVV7" : {[]string{".official-shop-brands", ".swiper-slide"}, 6},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "VV4" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 3},
    "V1" : {[]string{".span12"}, 0},
    "V2" : {[]string{".span6"}, 0},
    "V3" : {[]string{".span6"}, 1},
    "V4" : {[]string{".row-fluid"}, 2},
    "V5" : {[]string{".row-fluid"}, 3},
}

var SophisticatedApps = SophisticatedMobile

var ConvenienceDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".small-brand-logo"}, 0},
    "V2" : {[]string{".small-brand-logo"}, 1},
    "V3" : {[]string{".small-brand-logo"}, 2},
    "V4" : {[]string{".small-brand-logo"}, 3},
    "V5" : {[]string{".small-brand-logo"}, 4},
    "V6" : {[]string{".small-brand-logo"}, 5},
    "V7" : {[]string{".small-brand-logo"}, 6},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "E1" : {[]string{".official-promo-top-video"}, 0},
    "VVV1" : {[]string{".span6.official-shop-promo__type2", ".mt-20"}, 0},
    "VVV2" : {[]string{".span6.official-shop-promo__type2"}, 1},
    "B1" : {[]string{".official-shop-branding"}, 1},
}

var ConvenienceMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "V1" : {[]string{".official-shop-brands", ".swiper-slide"}, 0},
    "V2" : {[]string{".official-shop-brands", ".swiper-slide"}, 1},
    "V3" : {[]string{".official-shop-brands", ".swiper-slide"}, 2},
    "V4" : {[]string{".official-shop-brands", ".swiper-slide"}, 3},
    "V5" : {[]string{".official-shop-brands", ".swiper-slide"}, 4},
    "V6" : {[]string{".official-shop-brands", ".swiper-slide"}, 5},
    "V7" : {[]string{".official-shop-brands", ".swiper-slide"}, 6},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VVV1" : {[]string{".row-fluid.mt-20"}, 0},
    "VVV2" : {[]string{".row-fluid.mt-20"}, 1},
}

var ConvenienceApps = ConvenienceMobile

var BigPromoDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "T" : {[]string{".official-shop-top-video-half"}, 0},
    "V1" : {[]string{".official-shop-promo__type2"}, 0},
    "V2" : {[]string{".official-shop-promo__type2"}, 1},
    "V3" : {[]string{".official-shop-promo__type2"}, 2},
    "V4" : {[]string{".official-shop-promo__type2"}, 3},
}

var BigPromoMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "T" : {[]string{".official-shop-video"}, 0},
    "V1" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 0},
    "V2" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 1},
    "V3" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 2},
    "V4" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 3},
}

var BigPromoApps = BigPromoMobile

var FuturisticDesktop = map[string]ScrapeElement {
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "V1" : {[]string{".span3"}, 0},
    "V2" : {[]string{".span3"}, 1},
    "V3" : {[]string{".span3"}, 2},
    "V4" : {[]string{".span3"}, 3},
    "B1" : {[]string{".official-shop-branding"}, 3},
}

var FuturisticMobile = map[string]ScrapeElement {
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "V1" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 0},
    "V2" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 1},
    "V3" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 2},
    "V4" : {[]string{".section-official-shop-promo__type3", ".swiper-slide"}, 3},
}

var FuturisticApps = FuturisticMobile

var ClassyDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".small-brand-logo"}, 0},
    "V2" : {[]string{".small-brand-logo"}, 1},
    "V3" : {[]string{".small-brand-logo"}, 2},
    "VV1" : {[]string{".official-banner__"}, 0},
    "VV2" : {[]string{".official-banner__"}, 1},
    "VV3" : {[]string{".official-banner__"}, 2},
    "VVV1" : {[]string{".swiper-slide"}, 0},
    "VVV2" : {[]string{".swiper-slide"}, 1},
    "VVV3" : {[]string{".swiper-slide"}, 2},
    "VVVV1" : {[]string{".span6.official-banner"}, 0},
    "VVVV2" : {[]string{".span6.official-banner"}, 1},
}

var ClassyMobile = map[string]ScrapeElement {
    "G" : {[]string{".section-title"}, 0},
    "VV1" : {[]string{".official-banner__"}, 0},
    "VV2" : {[]string{".official-banner__"}, 1},
    "VV3" : {[]string{".official-banner__"}, 2},
    "VVV1" : {[]string{".swiper-slide"}, 0},
    "VVV2" : {[]string{".swiper-slide"}, 1},
    "VVV3" : {[]string{".swiper-slide"}, 2},
    "VVVV1" : {[]string{".span6.official-banner"}, 0},
    "VVVV2" : {[]string{".span6.official-banner"}, 1},
}

var ClassyApps = ClassyMobile

var HomeApplianceDesktop = map[string]ScrapeElement {
    "VV1" : {[]string{".promo-banner-col3.hovered"}, 0},
    "VV2" : {[]string{".promo-banner-col3.hovered"}, 1},
    "VVV1" : {[]string{".promo-banner-col3", ".row-fluid.hovered"}, 0},
    "VVV2" : {[]string{".promo-banner-col3", ".philips-promo-small", ".span6.hovered"}, 0},
    "VVV3" : {[]string{".promo-banner-col3", ".philips-promo-small", ".span6.hovered"}, 1},
    "VVVV1" : {[]string{".small-brand-logo"}, 0},
    "VVVV2" : {[]string{".small-brand-logo"}, 1},
    "VVVV3" : {[]string{".small-brand-logo"}, 2},
    "VVVV4" : {[]string{".small-brand-logo"}, 3},
    "VVVV5" : {[]string{".small-brand-logo"}, 4},
}

var HomeApplianceMobile = map[string]ScrapeElement {
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "VVV1" : {[]string{".row-fluid"}, 0},
    "VVV2" : {[]string{".row-fluid", ".span6"}, 0},
    "VVV3" : {[]string{".row-fluid", ".span6"}, 1},
    "VVVV1" : {[]string{".official-carousel", ".swiper-slide"}, 0},
    "VVVV2" : {[]string{".official-carousel", ".swiper-slide"}, 1},
    "VVVV3" : {[]string{".official-carousel", ".swiper-slide"}, 2},
    "VVVV4" : {[]string{".official-carousel", ".swiper-slide"}, 3},
    "VVVV5" : {[]string{".official-carousel", ".swiper-slide"}, 4},
}

var HomeApplianceApps = HomeApplianceMobile

var AllInOneDesktop = map[string]ScrapeElement {
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "T1" : {[]string{".tab-title"}, 0},
    "T2" : {[]string{".tab-title"}, 1},
    "T3" : {[]string{".tab-title"}, 2},
    "T4" : {[]string{".tab-title"}, 3},
    "V1" : {[]string{".branding-banner"}, 0},
    "L1" : {[]string{"#tab-beauty", ".brand-list-box"}, 0},
    "L2" : {[]string{"#tab-beauty", ".brand-list-box"}, 1},
    "L3" : {[]string{"#tab-beauty", ".brand-list-box"}, 2},
    "L4" : {[]string{"#tab-beauty", ".brand-list-box"}, 3},
    "L5" : {[]string{"#tab-beauty", ".brand-list-box"}, 4},
    "L6" : {[]string{"#tab-beauty", ".brand-list-box"}, 5},
    "L7" : {[]string{"#tab-beauty", ".brand-list-box"}, 6},
    "L8" : {[]string{"#tab-beauty", ".brand-list-box"}, 7},
    "L9" : {[]string{"#tab-beauty", ".brand-list-box"}, 9},
    "L10" : {[]string{"#tab-beauty", ".brand-list-box"}, 10},
    "L11" : {[]string{"#tab-beauty", ".brand-list-box"}, 11},
    "L12" : {[]string{"#tab-beauty", ".brand-list-box"}, 12},
    "L13" : {[]string{"#tab-beauty", ".brand-list-box"}, 13},
    "V2" : {[]string{".branding-banner"}, 1},
    "LL1" : {[]string{"#tab-baby", ".brand-list-box"}, 0},
    "LL2" : {[]string{"#tab-baby", ".brand-list-box"}, 1},
    "LL3" : {[]string{"#tab-baby", ".brand-list-box"}, 2},
    "V3" : {[]string{".branding-banner"}, 2},
    "LLL1" : {[]string{"#tab-home", ".brand-list-box"}, 0},
    "LLL2" : {[]string{"#tab-home", ".brand-list-box"}, 1},
    "LLL3" : {[]string{"#tab-home", ".brand-list-box"}, 2},
    "LLL4" : {[]string{"#tab-home", ".brand-list-box"}, 3},
    "LLL5" : {[]string{"#tab-home", ".brand-list-box"}, 4},
    "LLL6" : {[]string{"#tab-home", ".brand-list-box"}, 5},
    "LLL7" : {[]string{"#tab-home", ".brand-list-box"}, 6},
    "LLL8" : {[]string{"#tab-home", ".brand-list-box"}, 7},
    "V4" : {[]string{".branding-banner"}, 3},
    "LLLL1" : {[]string{"#tab-food", ".brand-list-box"}, 0},
    "LLLL2" : {[]string{"#tab-food", ".brand-list-box"}, 1},
    "LLLL3" : {[]string{"#tab-food", ".brand-list-box"}, 2},
    "VVV1" : {[]string{".official-shop-promo__banner-half"}, 0},
    "VVV2" : {[]string{".official-shop-promo__banner-half"}, 1},
}

var AllInOneMobile = map[string]ScrapeElement {
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "T1" : {[]string{".category-tab-link"}, 0},
    "T2" : {[]string{".category-tab-link"}, 1},
    "T3" : {[]string{".category-tab-link"}, 2},
    "T4" : {[]string{".category-tab-link"}, 3},
    "L1" : {[]string{"#tab-beauty", ".brand-list-box"}, 0},
    "L2" : {[]string{"#tab-beauty", ".brand-list-box"}, 1},
    "L3" : {[]string{"#tab-beauty", ".brand-list-box"}, 2},
    "L4" : {[]string{"#tab-beauty", ".brand-list-box"}, 3},
    "L5" : {[]string{"#tab-beauty", ".brand-list-box"}, 4},
    "L6" : {[]string{"#tab-beauty", ".brand-list-box"}, 5},
    "L7" : {[]string{"#tab-beauty", ".brand-list-box"}, 6},
    "L8" : {[]string{"#tab-beauty", ".brand-list-box"}, 7},
    "L9" : {[]string{"#tab-beauty", ".brand-list-box"}, 9},
    "L10" : {[]string{"#tab-beauty", ".brand-list-box"}, 10},
    "L11" : {[]string{"#tab-beauty", ".brand-list-box"}, 11},
    "L12" : {[]string{"#tab-beauty", ".brand-list-box"}, 12},
    "L13" : {[]string{"#tab-beauty", ".brand-list-box"}, 13},
    "LL1" : {[]string{"#tab-baby", ".brand-list-box"}, 0},
    "LL2" : {[]string{"#tab-baby", ".brand-list-box"}, 1},
    "LL3" : {[]string{"#tab-baby", ".brand-list-box"}, 2},
    "LLL1" : {[]string{"#tab-home", ".brand-list-box"}, 0},
    "LLL2" : {[]string{"#tab-home", ".brand-list-box"}, 1},
    "LLL3" : {[]string{"#tab-home", ".brand-list-box"}, 2},
    "LLL4" : {[]string{"#tab-home", ".brand-list-box"}, 3},
    "LLL5" : {[]string{"#tab-home", ".brand-list-box"}, 4},
    "LLL6" : {[]string{"#tab-home", ".brand-list-box"}, 5},
    "LLL7" : {[]string{"#tab-home", ".brand-list-box"}, 6},
    "LLL8" : {[]string{"#tab-home", ".brand-list-box"}, 7},
    "LLLL1" : {[]string{"#tab-food", ".brand-list-box"}, 0},
    "LLLL2" : {[]string{"#tab-food", ".brand-list-box"}, 1},
    "LLLL3" : {[]string{"#tab-food", ".brand-list-box"}, 2},
    "VVV1" : {[]string{".span6"}, 0},
    "VVV2" : {[]string{".span6"}, 1},
}

var AllInOneApps = AllInOneMobile

var Scrapers = map[string]interface{} {
    "advanced-Desktop" : AdvancedDesktop,
    "advanced-Mobile" : AdvancedMobile,
    "advanced-Apps" : AdvancedApps,

    "fashion-Desktop" : FashionDesktop,
    "fashion-Mobile" : FashionMobile,
    "fashion-Apps" : FashionApps,

    "sophisticated-Desktop" : SophisticatedDesktop,
    "sophisticated-Mobile" : SophisticatedMobile,
    "sophisticated-Apps" : SophisticatedApps,

    "convenience-Desktop" : ConvenienceDesktop,
    "convenience-Mobile" : ConvenienceMobile,
    "convenience-Apps" : ConvenienceApps,

    "big-promo-Desktop" : BigPromoDesktop,
    "big-promo-Mobile" : BigPromoMobile,
    "big-promo-Apps" : BigPromoApps,

    "futuristic-Desktop" : FuturisticDesktop,
    "futuristic-Mobile" : FuturisticMobile,
    "futuristic-Apps" : FuturisticApps,

    "classy-Desktop" : ClassyDesktop,
    "classy-Mobile" : ClassyMobile,
    "classy-Apps" : ClassyApps,

    "home-appliance-Desktop" : HomeApplianceDesktop,
    "home-appliance-Mobile" : HomeApplianceMobile,
    "home-appliance-Apps" : HomeApplianceApps,

    "all-in-one-Desktop" : AllInOneDesktop,
    "all-in-one-Mobile" : AllInOneMobile,
    "all-in-one-Apps" : AllInOneApps,
}