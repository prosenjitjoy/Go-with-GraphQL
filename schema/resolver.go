package schema

import (
	"context"
	"main/database"
	"main/model"

	"github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

var (
	db                 = database.DBinstance()
	productCollection  = database.OpenCollection(db, "products")
	categoryCollection = database.OpenCollection(db, "categories")
	reviewCollection   = database.OpenCollection(db, "reviews")
)

type Resolver struct {
	Query
	Mutation
}

type Query struct{}

func (*Query) Hello() string {
	return "Hello World!"
}

func (*Query) Number() int32 {
	return 27
}

func (*Query) Float() float64 {
	return 2.7
}

func (*Query) IsCool() bool {
	return true
}

func (*Query) List() []string {
	return []string{"Hello", "GraphQL"}
}

func (*Query) Categories(ctx context.Context) ([]model.Category, error) {
	database.OpenCollection(db, "categories")
	query := "FOR item IN categories RETURN item"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	categories := []model.Category{}
	for {
		var category model.Category
		_, err := cursor.ReadDocument(ctx, &category)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (*Query) Products(ctx context.Context) ([]model.Product, error) {
	database.OpenCollection(db, "products")
	query := "FOR item IN products RETURN item"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	products := []model.Product{}
	for {
		var product model.Product
		_, err := cursor.ReadDocument(ctx, &product)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (*Query) Reviews(ctx context.Context) ([]model.Review, error) {
	database.OpenCollection(db, "reviews")
	query := "FOR item IN reviews RETURN item"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	reviews := []model.Review{}
	for {
		var review model.Review
		_, err := cursor.ReadDocument(ctx, &review)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (*Query) Category(ctx context.Context, args *struct{ ID string }) (model.Category, error) {
	categoryID := args.ID

	var category model.Category
	_, err := categoryCollection.ReadDocument(context.TODO(), categoryID, &category)
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (*Query) Product(ctx context.Context, args *struct{ ID string }) (model.Product, error) {
	productID := args.ID

	var product model.Product
	_, err := productCollection.ReadDocument(context.TODO(), productID, &product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (*Query) Review(ctx context.Context, args *struct{ ID string }) (model.Review, error) {
	reviewID := args.ID

	var review model.Review
	_, err := reviewCollection.ReadDocument(context.TODO(), reviewID, &review)
	if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

type Mutation struct{}

func (*Mutation) AddCategory(ctx context.Context, args *struct{ Input model.CategoryInput }) (model.Category, error) {
	category := model.CategoryInput{
		Name: args.Input.Name,
	}

	newCategory := model.Category{
		ID:            uuid.NewString(),
		CategoryInput: category,
	}

	_, err := categoryCollection.CreateDocument(ctx, newCategory)
	if err != nil {
		return model.Category{}, err
	}

	return newCategory, nil
}

func (*Mutation) AddProduct(ctx context.Context, args *struct{ Input model.ProductInput }) (model.Product, error) {
	product := model.ProductInput{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		Quantity:    args.Input.Quantity,
		Price:       args.Input.Price,
		Image:       args.Input.Image,
		OnSale:      args.Input.OnSale,
		CategoryID:  args.Input.CategoryID,
	}

	newProduct := model.Product{
		ID:           uuid.NewString(),
		ProductInput: product,
	}

	_, err := productCollection.CreateDocument(ctx, newProduct)
	if err != nil {
		return model.Product{}, err
	}

	return newProduct, nil
}

func (*Mutation) AddReview(ctx context.Context, args *struct{ Input model.ReviewInput }) (model.Review, error) {
	review := model.ReviewInput{
		Date:      args.Input.Date,
		Title:     args.Input.Title,
		Comment:   args.Input.Comment,
		Rating:    args.Input.Rating,
		ProductID: args.Input.ProductID,
	}

	newReview := model.Review{
		ID:          uuid.NewString(),
		ReviewInput: review,
	}

	_, err := reviewCollection.CreateDocument(ctx, newReview)
	if err != nil {
		return model.Review{}, err
	}

	return newReview, nil
}

func (*Mutation) DeleteCategory(ctx context.Context, args *struct{ ID string }) (string, error) {
	categoryID := args.ID

	meta, err := categoryCollection.RemoveDocument(ctx, categoryID)
	if err != nil {
		return "", err
	}

	return meta.Key, err
}

func (*Mutation) DeleteProduct(ctx context.Context, args *struct{ ID string }) (string, error) {
	productID := args.ID

	meta, err := productCollection.RemoveDocument(ctx, productID)
	if err != nil {
		return "", err
	}

	return meta.Key, err
}

func (*Mutation) DeleteReview(ctx context.Context, args *struct{ ID string }) (string, error) {
	reviewID := args.ID

	meta, err := reviewCollection.RemoveDocument(ctx, reviewID)
	if err != nil {
		return "", err
	}

	return meta.Key, err
}

func (*Mutation) UpdateCategory(ctx context.Context, args *struct {
	ID    string
	Input model.CategoryInput
}) (string, error) {
	meta, err := categoryCollection.UpdateDocument(ctx, args.ID, args.Input)
	if err != nil {
		return "", err
	}

	return meta.Key, nil
}

func (*Mutation) UpdateProduct(ctx context.Context, args *struct {
	ID    string
	Input model.ProductInput
}) (string, error) {
	meta, err := productCollection.UpdateDocument(ctx, args.ID, args.Input)
	if err != nil {
		return "", err
	}

	return meta.Key, nil
}

func (*Mutation) UpdateReview(ctx context.Context, args *struct {
	ID    string
	Input model.ReviewInput
}) (string, error) {
	meta, err := reviewCollection.UpdateDocument(ctx, args.ID, args.Input)
	if err != nil {
		return "", err
	}

	return meta.Key, nil
}
