package account

import (
	"testing"
	"time"
)

func TestRateLimiterDefault(t *testing.T) {
	rl := NewRateLimiter(0)
	if rl == nil {
		t.Fatal("expected rate limiter to not be nil")
	}
}

func TestRateLimiterCustomInterval(t *testing.T) {
	rl := NewRateLimiter(100)
	if rl == nil {
		t.Fatal("expected rate limiter to not be nil")
	}
}

func TestRateLimiterUpdateInterval(t *testing.T) {
	rl := NewRateLimiter(100)
	rl.UpdateInterval(200)
	// 这只是一个简单的验证，确保 UpdateInterval 不会出错
}

func TestRateLimiterWait(t *testing.T) {
	rl := NewRateLimiter(50) // 50ms 的间隔
	
	start := time.Now()
	rl.Wait("test-account")
	elapsed := time.Since(start)
	// 第一次等待不应该有延迟
	if elapsed > 10*time.Millisecond {
		t.Fatalf("first wait should not have significant delay, got %s", elapsed)
	}
	
	// 第二次等待应该有延迟
	start = time.Now()
	rl.Wait("test-account")
	elapsed = time.Since(start)
	if elapsed < 40*time.Millisecond {
		t.Fatalf("second wait should have delay, got %s", elapsed)
	}
}

func TestRateLimiterDifferentAccounts(t *testing.T) {
	rl := NewRateLimiter(100)
	
	// 对于不同的账号，不应该有等待
	start := time.Now()
	rl.Wait("account-1")
	elapsed1 := time.Since(start)
	
	start = time.Now()
	rl.Wait("account-2")
	elapsed2 := time.Since(start)
	
	if elapsed1 > 10*time.Millisecond || elapsed2 > 10*time.Millisecond {
		t.Fatalf("different accounts should not cause delay, got %s and %s", elapsed1, elapsed2)
	}
}
