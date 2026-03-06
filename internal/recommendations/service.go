package recommendations

import "user-profile-system-backend-go/internal/models"

func RecommendForUser(userID string, history []models.UserHistory, favorites []models.UserFavorite, allContent []models.ContentItem) []models.ContentItem {

    profile := BuildUserProfile(history, favorites)

    scored := make([]struct{
        item  models.ContentItem
        score float64
    }, 0)

    for _, c := range allContent {
        score := FinalScore(profile, c)
        scored = append(scored, struct{
            item  models.ContentItem
            score float64
        }{
            item: c,
            score: score,
        })
    }

    // Sort by descending score
    sort.Slice(scored, func(i, j int) bool {
        return scored[i].score > scored[j].score
    })

    // Return top 10
    out := []models.ContentItem{}
    for i := 0; i < len(scored) && i < 10; i++ {
        out = append(out, scored[i].item)
    }

    return out
}
