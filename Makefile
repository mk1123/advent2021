day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp template calendar/day-0$(day)/main.go; \
  	else \
  		mkdir calendar/day-$(day); \
		cp template calendar/day-$(day)/main.go; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/

