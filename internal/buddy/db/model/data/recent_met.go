package data

type RecentProfile struct {
	ID      string
	AddTime int64
}

func (bq *BuddyQueue) AddRecentProfiles() {

}

func (bq *BuddyQueue) DeleteRecentProfiles(ids ...string) {
	for _, v := range ids {
		for i := 0; i < len(bq.RecentMet); i++ {
			if bq.RecentMet[i].ID == v {
				bq.RecentMet = append(bq.RecentMet[:i], bq.RecentMet[i+1:]...)
				i--
			}
		}
	}
}

func (bq *BuddyQueue) removeDuplicatesInOrder(profiles []*RecentProfile) []*RecentProfile {
	if len(profiles) <= 0 {
		return nil
	}
	checks := make(map[string]bool)
	result := make([]*RecentProfile, 0)
	for _, v := range profiles {
		if ok := checks[v.ID]; !ok {
			checks[v.ID] = true
			result = append(result, v)
		}
	}
	return result
}
