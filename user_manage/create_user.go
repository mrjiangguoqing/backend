package models

import (
    "gorm.io/gorm"
    "time"
)



type User struct {
    ID           uint      `gorm:"primaryKey;autoIncrement"` // 主键，自增
    UserID       int       `gorm:"column:user_id;unique"`    // 用户 ID，唯一
    Email        string    `gorm:"column:email;unique;not null"` // 邮箱，唯一，非空
    PasswordHash string    `gorm:"column:password_hash;not null"` // 密码哈希，非空
    CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"` // 注册时间
    UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime"` // 更新时间
}





package services

import (
    "your_project/models"
    "gorm.io/gorm"
    "errors"
)

type CreateUserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Username string `json:"username"` // 如果需要用户名，可以添加到请求中
}

func CreateUser(db *gorm.DB, req CreateUserRequest) error {
    // 检查邮箱是否已存在
    var existingUser models.User
    if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        return errors.New("email already exists")
    }

    // 对密码进行哈希处理（假设你已经有一个哈希函数）
    passwordHash, err := HashPassword(req.Password)
    if err != nil {
        return err
    }

    // 创建新用户
    newUser := models.User{
        Email:        req.Email,
        PasswordHash: passwordHash,
        // 如果需要用户名，可以在这里赋值
    }

    // 插入数据库
    if err := db.Create(&newUser).Error; err != nil {
        return err
    }

    return nil
}

// 假设的密码哈希函数
func HashPassword(password string) (string, error) {
    // 这里可以使用 bcrypt 或其他哈希算法
    // 例如：return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return password, nil // 实际使用时请替换为真正的哈希逻辑
}



func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

	