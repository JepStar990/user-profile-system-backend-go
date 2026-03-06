package recommendations

import "user-profile-system-backend-go/internal/models"

func ContentSimilarity(profile UserPreferenceProfile, item models.ContentItem) float64 {
    score := 0.0

    // category match
    for _, cat := range profile.TopCategories {
        if item.Category == cat {
            score += 0.3
        }
    }

    // tags overlap
    for _, t := range profile.FrequentTags {
        for _, it := range item.Tags {
            if t == it {
                score += 0.2
            }
        }
    }

    // duration preference
    // short: < 300s, medium: < 900s, long otherwise
    // match → +0.1

    return score
}
