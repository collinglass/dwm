typedef union {
	float coords[3];
	char about[20];
} assocdata;

typedef struct {
	char *descript;
	int un_type;
	assocdata alsostuff;
} leaf;

int main () {
	leaf oak[3];
	int i;

	printf("Hello world!\n");

	for (i=0; i<3; i++) {
		oak[i].descript = "A greeting";
		oak[i].un_type = 1;
		oak[i].alsostuff.coords[0] = 3.14;
	}
	oak[2].alsostuff.coords[0] = 3.14;

	for (i = 0; i < 3; i++)
	{
		printf("%s\n", oak[i].descript);
		printf("%5.2f\n", oak[i].alsostuff.coords[0]);
	}
}