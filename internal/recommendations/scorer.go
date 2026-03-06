package recommendations

import "user-profile-system-backend-go/internal/models"

func FinalScore(profile UserPreferenceProfile, item models.ContentItem) float64 {
    behavioral := 0.0
    // Compute based on:
    // - how often user listened to similar items
    // - if user completed similar items
    // - if item matches user's active hours habit

    similarity := ContentSimilarity(profile, item)

    // Weighted final score
    return (behavioral * 0.6) + (similarity * 0.4)
}
