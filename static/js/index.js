$(() => {
    // Modal close buttons
    $(".modal button.delete, .modal button.modal-close, .modal .modal-background").click((e) => {
        e.preventDefault();
        $(e.target).closest('.modal').removeClass("is-active");
    })

    // Open Wizard Button
    $(".btn-openWizard").click((e) => {
        e.preventDefault();
        $(".modal.is-first").toggleClass("is-active");
    });

    // calendar guide buttons
    $(".calendar-service-buttons a.button").click((e) => {
        e.preventDefault();
        let serviceName = e.target.dataset.serviceName || $(e.target).closest('.button').data('serviceName');
        $(".modal.is-first").removeClass("is-active");
        $(`.modal.is-guide.is-${serviceName}-guide`).toggleClass("is-active");
    });

    // ics guide link
    $(".ics-link").click(e => {
        e.preventDefault();
        $(".modal.is-first").removeClass("is-active");
        $('.modal.is-ics-guide').toggleClass('is-active');
    })

    // faq button
    $('.button.btn-faq').click(e => {
        e.preventDefault();
        $('.modal.is-faq').toggleClass('is-active');
    })
});